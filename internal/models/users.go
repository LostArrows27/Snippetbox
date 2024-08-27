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

type UserData struct {
	ID   int
	Name string
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

func (m *UserModel) Authenticate(email, password string) (UserData, error) {
	// 1. check if email is exist
	var hasedPassword []byte
	var user UserData = UserData{}

	queryStr := "SELECT id, hashed_password, name FROM users WHERE email = $1"

	err := m.DB.QueryRow(queryStr, email).Scan(&user.ID, &hasedPassword, &user.Name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, ErrInvalidCredentials
		} else {
			return user, err
		}
	}

	// 2. check if password match in database
	err = bcrypt.CompareHashAndPassword(hasedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return user, ErrInvalidCredentials
		} else {
			return user, err
		}
	}

	// 3. return UserID -> password correct
	return user, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
