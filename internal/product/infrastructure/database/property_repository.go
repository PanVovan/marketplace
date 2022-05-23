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

type ProductPropertyRepositoryPostgres struct {
	db     *sql.DB
	dao    *dao.ProductDao
	mapper *mapper.ProductPropertyMapper
}

func NewProductPropertyRepositoryPostgres(
	db *sql.DB,
) *ProductPropertyRepositoryPostgres {
	return &ProductPropertyRepositoryPostgres{
		db:     db,
		dao:    dao.NewProductDao(db),
		mapper: mapper.NewProductPropertyMapper(),
	}
}

func (pp *ProductPropertyRepositoryPostgres) Create(ctx context.Context, productID uuid.UUID, property model.CreateProductPropertyParams) (uuid.UUID, error) {
	return pp.dao.CreateProductProperty(ctx, sqlc.CreateProductPropertyParams{
		ProductID: productID,
		Name:      property.Name,
		Value:     property.Value,
	})
}
func (pp *ProductPropertyRepositoryPostgres) GetOne(ctx context.Context, propertyID uuid.UUID) (model.ProductProperty, error) {
	dto, err := pp.dao.GetProductPropertyByID(ctx, propertyID)
	if err != nil {
		return model.ProductProperty{}, err
	}
	return pp.mapper.FromDTOToEntity(dto), nil
}

func (pp *ProductPropertyRepositoryPostgres) GetAll(ctx context.Context, productID uuid.UUID) ([]model.ProductProperty, error) {
	dtos, err := pp.dao.GetProductPropertiesByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	entities := make([]model.ProductProperty, 0)
	for _, dto := range dtos {
		entities = append(entities, pp.mapper.FromDTOToEntity(dto))
	}
	return entities, nil
}

func (pp *ProductPropertyRepositoryPostgres) Update(ctx context.Context, propertyID uuid.UUID, params model.UpdateProductPropertyParams) error {
	return pp.dao.UpdateProductProperty(ctx, propertyID, dao.UpdateProductPropertyParams{
		Name:  &params.Name,
		Value: &params.Value,
	})

}
func (pp *ProductPropertyRepositoryPostgres) Delete(ctx context.Context, propertyID uuid.UUID) error {
	return pp.dao.DeleteProductProperty(ctx, propertyID)
}
