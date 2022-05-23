package database

import (
	"context"
	"database/sql"
	"makretplace/internal/product/domain/mapper"
	"makretplace/internal/product/domain/model"
	"makretplace/internal/product/infrastructure/database/dao"
	"makretplace/internal/product/infrastructure/database/sqlc"

	"github.com/google/uuid"
)

type ProductImageRepositoryPostgres struct {
	db     *sql.DB
	dao    *dao.ProductDao
	mapper *mapper.ProductImageMapper
}

func NewProductImageRepositoryPostgres(db *sql.DB, dao *dao.ProductDao) *ProductImageRepositoryPostgres {
	return &ProductImageRepositoryPostgres{
		db:     db,
		dao:    dao,
		mapper: mapper.NewProductImageMapper(),
	}
}

func (pi *ProductImageRepositoryPostgres) Create(ctx context.Context, productID uuid.UUID, filename string) (uuid.UUID, error) {
	return pi.dao.CreateImage(ctx, sqlc.CreateImageParams{
		ProductID: productID,
		File:      filename,
	})
}

func (pi *ProductImageRepositoryPostgres) GetOne(ctx context.Context, imageID uuid.UUID) (model.ProductImage, error) {
	image, err := pi.dao.GetImageByID(ctx, imageID)
	if err != nil {
		return model.ProductImage{}, err
	}
	return pi.mapper.FromDTOToEntity(image), nil
}

func (pi *ProductImageRepositoryPostgres) GetAll(ctx context.Context, productID uuid.UUID) ([]model.ProductImage, error) {
	dtos, err := pi.dao.GetImagesByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	entities := make([]model.ProductImage, 0)
	for _, dto := range dtos {
		entities = append(entities, pi.mapper.FromDTOToEntity(dto))
	}
	return entities, nil
}
func (pi *ProductImageRepositoryPostgres) Update(ctx context.Context, imageID uuid.UUID, filename string) error {
	return pi.dao.UpdateProductImage(ctx, filename, imageID)
}
func (pi *ProductImageRepositoryPostgres) Delete(ctx context.Context, imageID uuid.UUID) error {
	return pi.dao.DeleteImage(ctx, imageID)
}
