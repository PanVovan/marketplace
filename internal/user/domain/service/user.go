package service

import (
	"github.com/google/uuid"
)

type UserService interface {
	GenerateToken(userId uuid.UUID) (string, error)
	ParseToken(accessToken string) (uuid.UUID, error)
	GeneratePasswordHash(password string) string
}
