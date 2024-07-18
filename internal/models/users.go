package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/lib/pq" // Import the PostgreSQL driver package
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created) VALUES($1, $2, $3, NOW())`

	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))

	if err != nil {
		var postgresError *pq.Error

		if errors.As(err, &postgresError) {
			if postgresError.Code == "23505" && strings.Contains(postgresError.Message, "users_uc_email") {
				return ErrDuplicateEmail
			}
		}
		return err

	}

	return nil

}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	// 1. check if email is exist
	var id int
	var hasedPassword []byte

	queryStr := "SELECT id, hashed_password FROM users WHERE email = $1"

	err := m.DB.QueryRow(queryStr, email).Scan(&id, &hasedPassword)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	// 2. check if password match in database
	err = bcrypt.CompareHashAndPassword(hasedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	// 3. return UserID -> password correct
	return id, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
