package model

import "github.com/google/uuid"

type Seller struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password string
}

type CreateSellerParams struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type UpdateSellerParams struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Name     *string `json:"name"`
}
