package service

import (
	"github.com/google/uuid"
)

type SellerService interface {
	GenerateToken(SellerId uuid.UUID) (string, error)
	ParseToken(accessToken string) (uuid.UUID, error)
	GeneratePasswordHash(password string) string
}
