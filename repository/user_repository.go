package repository

import (
	"log"

	"github.com/elhaqeeem/my-golang-project/db"
	"github.com/elhaqeeem/my-golang-project/service"
)

// GetUsers untuk mendapatkan semua user
func GetUsers() ([]service.User, error) {
	rows, err := db.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var users []service.User
	for rows.Next() {
		var user service.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Fatal("Error scanning row:", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUserByID untuk mendapatkan user berdasarkan ID
func GetUserByID(id string) (service.User, error) {
	var user service.User
	err := db.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

// CreateUser untuk menambahkan user baru
func CreateUser(user service.User) (service.User, error) {
	err := db.DB.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return user, err
	}
	return user, nil
}
