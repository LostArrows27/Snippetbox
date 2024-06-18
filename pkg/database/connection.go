package database

import (
	"database/sql"

	_ "github.com/lib/pq" // Import the PostgreSQL driver package
)

func OpenDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
