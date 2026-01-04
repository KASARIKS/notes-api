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

type NotesStore interface {
	CreateNote(userID int, note Note) error
	GetNotesByUserId(userID int) ([]Note, error)
}

type Note struct {
	Id        int    `json:"id"`
	UserID    int    `json:"userId"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
}

type NotePayload struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
