package service

import (
	"github.com/elhaqeeem/my-golang-project/model"
	"github.com/elhaqeeem/my-golang-project/repository"
)

// GetAllUsers untuk mendapatkan semua user
func GetAllUsers() ([]model.User, error) {
	return repository.GetUsers()
}

// GetUserByID untuk mendapatkan user berdasarkan ID
func GetUserByID(id string) (model.User, error) {
	return repository.GetUserByID(id)
}

// CreateUser untuk membuat user baru
func CreateUser(user model.User) (model.User, error) {
	return repository.CreateUser(user)
}
