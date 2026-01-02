package types

type UserStore interface {
	GetUserById(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(User) error
	DeleteUserById(id int) error
}

type User struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type RegisterUserPayload struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DeleteUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
