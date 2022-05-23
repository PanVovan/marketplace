package mapper

import (
	"makretplace/internal/brand/domain/model"
	"makretplace/internal/brand/infrastructure/database/sqlc"
)

type BrandMapper struct{}

func NewBrandMapper() *BrandMapper {
	return &BrandMapper{}
}

func (b *BrandMapper) FromDTOToEntity(dto sqlc.Brand) model.Brand {
	return model.Brand{
		ID:   dto.ID,
		Name: dto.Name,
	}
}

func (b *BrandMapper) FromEntityToDTO(model model.Brand) sqlc.Brand {
	return sqlc.Brand{
		ID:   model.ID,
		Name: model.Name,
	}
}
