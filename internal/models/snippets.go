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

	// 1. query rows
	query := `SELECT id, title, content, created, expires FROM snippets
			  WHERE expires > CURRENT_TIMESTAMP 
			  ORDER BY id DESC 
			  LIMIT 10`

	rows, err := m.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 2. loop rows
	snippetsList := []*Snippet{}

	for rows.Next() {
		snippet := &Snippet{}

		err := rows.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)

		if err != nil {
			return nil, err
		}

		snippetsList = append(snippetsList, snippet)

	}

	// 3. check loop error
	if err := rows.Err(); err != nil {
		return nil, err

	}

	return snippetsList, nil

}
