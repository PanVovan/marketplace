package database

import (
	"context"
	"database/sql"
	"makretplace/internal/category/domain/mapper"
	"makretplace/internal/category/domain/model"
	"makretplace/internal/category/infrastructure/database/dao"
	"makretplace/internal/category/infrastructure/database/sqlc"

	"github.com/google/uuid"
)

type CategoryRepositoryPostgres struct {
	db             *sql.DB
	categoryDao    *dao.CategoryDao
	categoryMapper *mapper.CategoryMapper
}

func NewCategoryRepositoryPostgres(db *sql.DB) *CategoryRepositoryPostgres {
	return &CategoryRepositoryPostgres{
		db:             db,
		categoryDao:    dao.NewCategoryDao(db),
		categoryMapper: mapper.NewCategoryMapper(),
	}
}

func (cr *CategoryRepositoryPostgres) GetAll(ctx context.Context, limit, page int32) ([]model.Category, error) {
	dtos, err := cr.categoryDao.GetCategories(ctx, sqlc.GetCategoriesParams{
		Limit:  limit,
		Offset: limit * (page - 1),
	})
	if err != nil {
		return nil, err
	}
	var entities []model.Category
	for _, dto := range dtos {
		entities = append(entities, cr.categoryMapper.FromDTOToEntity(dto))
	}

	return entities, nil

}

func (cr *CategoryRepositoryPostgres) Create(ctx context.Context, name string) (uuid.UUID, error) {
	return cr.categoryDao.CreateCategory(ctx, name)
}

func (cr *CategoryRepositoryPostgres) GetOne(ctx context.Context, categoryID uuid.UUID) (model.Category, error) {
	dto, err := cr.categoryDao.GetCategoryById(ctx, categoryID)
	if err != nil {
		return model.Category{}, err
	}

	return cr.categoryMapper.FromDTOToEntity(dto), nil
}

func (cr *CategoryRepositoryPostgres) Update(ctx context.Context, categoryID uuid.UUID, name string) error {
	return cr.categoryDao.UpdateCategory(ctx, categoryID, name)
}

func (cr *CategoryRepositoryPostgres) Delete(ctx context.Context, categoryID uuid.UUID) error {
	return cr.categoryDao.DeleteCategory(ctx, categoryID)
}
