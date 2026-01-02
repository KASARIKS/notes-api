package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kasariks/notes_api/config"
	"github.com/kasariks/notes_api/types"
	"github.com/kasariks/notes_api/utils"
)

const UserKey string = "userID"

func CreateJWT(secret []byte, userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiresAt": config.Envs.JWTExpirationInSeconds,
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

// Check token
func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get token
		tokenString := getTokenFromRequest(r)

		// Validate token
		token, err := validateToken(tokenString)
		if err != nil {
			log.Printf("failed to validate the token: %v", err)
			permissionDenied(w)
			return
		}

		// Send userID to handler through context
		claims := token.Claims.(jwt.MapClaims)
		str := claims[UserKey].(string)

		userID, err := strconv.Atoi(str)
		if err != nil {
			log.Printf("failed to convert user id: %v", err)
			permissionDenied(w)
			return
		}

		u, err := store.GetUserById(userID)
		if err != nil {
			log.Printf("failed to get the user by id: %v", err)
			permissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")

	return tokenAuth
}

func validateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected string method: %v", t.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserIdFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}

	return userID
}
