package mapper

import (
	"makretplace/internal/seller/domain/model"
	"makretplace/internal/seller/infrastructure/database/sqlc"
)

type SellerInfoMapper struct{}

func (s *SellerInfoMapper) FromDTOToEntity(dto sqlc.SellerInfo) model.SellerInfo {
	return model.SellerInfo{
		ID:   dto.ID,
		Name: dto.Name,
	}
}

func (s *SellerInfoMapper) FromEntityToDTO(model model.SellerInfo) sqlc.SellerInfo {
	return sqlc.SellerInfo{
		ID:   model.ID,
		Name: model.Name,
	}
}
