package mapper

import (
	"makretplace/internal/product/domain/model"
	"makretplace/internal/product/infrastructure/database/sqlc"
)

type productPropertyMapper struct{}

func (p *productPropertyMapper) FromDTOToEntity(dto sqlc.ProductProperty) model.ProductProperty {
	return model.ProductProperty{
		ID:    dto.ID,
		Name:  dto.Name,
		Value: dto.Value,
	}
}

func (p *productPropertyMapper) FromEntityToDTO(model model.ProductProperty) sqlc.ProductProperty {
	return sqlc.ProductProperty{
		ID:    model.ID,
		Name:  model.Name,
		Value: model.Value,
	}
}
