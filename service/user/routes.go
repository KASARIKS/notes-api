package user

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/kasariks/notes_api/config"
	"github.com/kasariks/notes_api/service/auth"
	"github.com/kasariks/notes_api/types"
	"github.com/kasariks/notes_api/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

// eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzQXQiOjYwNDgwMCwidXNlcklEIjoxfQ.UVAHARPkfa7D4Fe_lDzKGef5u7Ei7-hh-C-XeLpXnDpTIVFhypQPr1dnAB1AaWIy80lwsxUjh4oXAatg07RRtg

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /register", h.HandleRegister)
	router.HandleFunc("POST /login", h.HandleLogin)
	router.HandleFunc("DELETE /delete_user", auth.WithJWTAuth(h.HandleUserDeleting, h.store))
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	// Get json payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate email
	if _, err := mail.ParseAddress(payload.Email); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Check if the user with the email exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	// Create user
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		Nickname: payload.Nickname,
		Email:    payload.Email,
		Password: hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Get json payload
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate email
	if _, err := mail.ParseAddress(payload.Email); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Check if the user with the email exists
	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s doesn't exist", payload.Email))
		return
	}

	// Compare passwords
	err = auth.ComparePasswords(u.Password, payload.Password)
	if err == auth.IncorrectPassword {
		utils.WriteError(w, http.StatusNotAcceptable, err)
		return
	} else if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Send token
	token, err := auth.CreateJWT([]byte(config.Envs.JWTSecret), u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) HandleUserDeleting(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIdFromContext(r.Context())

	if err := h.store.DeleteUserById(userID); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}
