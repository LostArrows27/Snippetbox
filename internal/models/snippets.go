package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	id := 0

	query := `INSERT INTO snippets (title, content, created, expires)
			  VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '1 day' * $3) RETURNING id;`

	err := m.DB.QueryRow(query, title, content, expires).Scan(&id)

	if err != nil {
		return id, err
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {

	query := `SELECT id, title, content, created, expires FROM snippets
			  WHERE expires > CURRENT_TIMESTAMP AND id = $1`

	s := &Snippet{}

	err := m.DB.QueryRow(query, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
