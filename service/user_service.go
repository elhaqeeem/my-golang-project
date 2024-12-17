package service

import (
	"github.com/my-golang-project/repository"
)

// User struct untuk model user
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GetAllUsers untuk mendapatkan semua user
func GetAllUsers() ([]User, error) {
	return repository.GetUsers()
}

// GetUserByID untuk mendapatkan user berdasarkan ID
func GetUserByID(id string) (User, error) {
	return repository.GetUserByID(id)
}

// CreateUser untuk membuat user baru
func CreateUser(user User) (User, error) {
	return repository.CreateUser(user)
}
