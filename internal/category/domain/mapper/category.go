package mapper

import (
	"makretplace/internal/category/domain/model"
	"makretplace/internal/category/infrastructure/database/sqlc"
)

type CategoryMapper struct{}

func NewCategoryMapper() *CategoryMapper {
	return &CategoryMapper{}
}

func (c *CategoryMapper) FromDTOToEntity(dto sqlc.Category) model.Category {
	return model.Category{
		ID:   dto.ID,
		Name: dto.Name,
	}
}

func (c *CategoryMapper) FromEntityToDTO(model model.Category) sqlc.Category {
	return sqlc.Category{
		ID:   model.ID,
		Name: model.Name,
	}
}
