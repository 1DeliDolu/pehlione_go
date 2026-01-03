package models

import (
	"time"

	"pehlione.com/my_go_app/database"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetAllUsers returns all users from database
func GetAllUsers() ([]User, error) {
	rows, err := database.DB.Query("SELECT id, username, email, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// GetUserByID returns a user by ID
func GetUserByID(id int) (*User, error) {
	var u User
	err := database.DB.QueryRow("SELECT id, username, email, created_at FROM users WHERE id = ?", id).
		Scan(&u.ID, &u.Username, &u.Email, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// CreateUser inserts a new user into database
func CreateUser(username, email string) (int64, error) {
	result, err := database.DB.Exec("INSERT INTO users (username, email) VALUES (?, ?)", username, email)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// UpdateUser updates an existing user
func UpdateUser(id int, username, email string) error {
	_, err := database.DB.Exec("UPDATE users SET username = ?, email = ? WHERE id = ?", username, email, id)
	return err
}

// DeleteUser deletes a user by ID
func DeleteUser(id int) error {
	_, err := database.DB.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
