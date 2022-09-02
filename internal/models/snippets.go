package models

import (
	"database/sql"
	"errors"
	"time"
)

// Snippet hold the data for an individual snippet.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel manupulates by snippets
type SnippetModel struct {
	DB *sql.DB
}

// Insert stores a snippet to a DB.
func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get finds a snippet by specified id.
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	s := &Snippet{}

	stmt := `SELECT id, title, content, created, expires
	FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?`

	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):

		}

		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

// Latest returns 10 latestly created snippets.
func (m *SnippetModel) Latest() (*Snippet, error) {
	return nil, nil
}
