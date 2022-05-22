package application

import (
	"context"
	"makretplace/internal/seller/domain/model"
	"makretplace/internal/seller/domain/repository"
	"makretplace/internal/seller/domain/service"

	"github.com/google/uuid"
)

type SignUpUseCase struct {
	repository repository.SellerRepository
	service    service.SellerService
}

func NewSignUpUseCase(
	repository repository.SellerRepository,
	service service.SellerService,
) *SignUpUseCase {
	return &SignUpUseCase{
		repository: repository,
		service:    service,
	}
}

func (s *SignUpUseCase) Execute(ctx context.Context, seller model.CreateSellerParams) (string, uuid.UUID, error) {
	seller.Password = s.service.GeneratePasswordHash(seller.Password)
	sellerId, err := s.repository.Create(ctx, seller)
	if err != nil {
		return "", uuid.Nil, err
	}
	token, err := s.service.GenerateToken(sellerId)
	return token, sellerId, err
}
