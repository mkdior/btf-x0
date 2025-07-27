package models

import domain "github.com/mkdior/btf-x0/internal/models/user"

type User struct {
	ID      int `json:"id"`
	Balance int `json:"balance"`
}

type CreateUserRequest struct {
	Data []User `json:"data"`
}

// @TODO(Hamza) - remove this dep
func (u User) ToDomain() domain.User {
	return domain.User{ID: u.ID, Balance: u.Balance}
}
