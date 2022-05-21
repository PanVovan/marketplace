package mapper

import (
	"makretplace/internal/product/domain/model"
	"makretplace/internal/product/infrastructure/database/sqlc"
)

type ProductImageMapper struct{}

func (p *ProductImageMapper) FromDTOToEntity(dto sqlc.ProductImage) model.ProductImage {
	return model.ProductImage{
		ID:   dto.ID,
		File: dto.File,
	}
}

func (p *ProductImageMapper) FromEntityToDTO(model model.ProductImage) sqlc.ProductImage {
	return sqlc.ProductImage{
		ID:   model.ID,
		File: model.File,
	}
}
