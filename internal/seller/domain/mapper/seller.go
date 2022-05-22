package mapper

import (
	"makretplace/internal/seller/domain/model"
	"makretplace/internal/seller/infrastructure/database/sqlc"
)

type SellerMapper struct{}

func NewSellerMapper() *SellerMapper {
	return &SellerMapper{}
}

func (s *SellerMapper) FromDTOToEntity(dto sqlc.Seller) model.Seller {
	return model.Seller{
		ID:       dto.ID,
		Email:    dto.Email,
		Password: dto.Password,
		Name:     dto.Name,
	}
}

func (s *SellerMapper) FromEntityToDTO(model model.Seller) sqlc.Seller {
	return sqlc.Seller{
		ID:       model.ID,
		Email:    model.Email,
		Password: model.Password,
		Name:     model.Name,
	}
}
