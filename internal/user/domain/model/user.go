package model

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Name     string    `json:"name" binding:"required"`
}

type CreateUserParams struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type UpdateUserParams struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Name     *string `json:"name"`
}
