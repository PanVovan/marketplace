package application

import (
	"context"
	"errors"
	"makretplace/internal/seller/domain/repository"
	"makretplace/internal/seller/domain/service"
)

type LoginUseCase struct {
	repository repository.SellerRepository
	service    service.SellerService
}

func NewLoginUseCase(
	repository repository.SellerRepository,
	service service.SellerService,
) *LoginUseCase {
	return &LoginUseCase{
		repository: repository,
		service:    service,
	}
}

type LoginSellerParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginUseCase) Execute(ctx context.Context, seller LoginSellerParams) (string, error) {
	item, err := l.repository.GetByEmail(ctx, seller.Email)
	if err != nil {
		return "", err
	}

	if item.Password != l.service.GeneratePasswordHash(seller.Password) {
		return "", errors.New("password incorrest")
	}

	return l.service.GenerateToken(item.ID)
}
