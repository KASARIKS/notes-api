package user

import (
	"database/sql"

	"github.com/kasariks/notes_api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	user := &types.User{}

	row := s.db.QueryRow("SELECT * FROM users WHERE id = :id;", sql.Named("id", id))
	err := row.Scan(
		&user.ID,
		&user.Nickname,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	user := &types.User{}

	row := s.db.QueryRow("SELECT * FROM users WHERE email = :email;", sql.Named("email", email))
	err := row.Scan(
		&user.ID,
		&user.Nickname,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (nickname, email, password) VALUES (:nickname, :email, :password);",
		sql.Named("nickname", user.Nickname),
		sql.Named("email", user.Email),
		sql.Named("password", user.Password))
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteUserById(id int) error {
	_, err := s.db.Exec("DELETE FROM users WHERE id = :id", sql.Named("id", id))
	if err != nil {
		return err
	}

	return nil
}
