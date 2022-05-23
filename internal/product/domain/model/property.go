package model

import "github.com/google/uuid"

type ProductProperty struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Value string    `json:"value"`
}

type CreateProductPropertyParams struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type UpdateProductPropertyParams struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
