package note

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

func (s *Store) CreateNote(userID int, note types.Note) error {
	_, err := s.db.Exec("INSERT INTO notes (userId, name, value, createdAt) VALUES (:userId, :name, :value, :createdAt)",
		sql.Named("userId", userID),
		sql.Named("name", note.Name),
		sql.Named("value", note.Value),
		sql.Named("createdAt", note.CreatedAt))

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetNotesByUserId(userID int) ([]types.Note, error) {
	notes := []types.Note{}

	rows, err := s.db.Query("SELECT * FROM notes WHERE userId = :userId", sql.Named("userId", userID))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		n, err := scanFromRows(rows)
		if err != nil {
			return nil, err
		}

		notes = append(notes, *n)
	}

	return notes, nil
}

func scanFromRows(rows *sql.Rows) (*types.Note, error) {
	var note types.Note
	err := rows.Scan(
		&note.Id,
		&note.UserID,
		&note.Name,
		&note.Value,
		&note.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &note, nil
}
