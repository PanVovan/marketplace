package database

import (
	"context"
	"database/sql"
	"makretplace/internal/brand/domain/mapper"
	"makretplace/internal/brand/domain/model"
	"makretplace/internal/brand/infrastructure/database/dao"
	"makretplace/internal/brand/infrastructure/database/sqlc"

	"github.com/google/uuid"
)

type BrandRepositoryPostgres struct {
	db          *sql.DB
	brandDao    *dao.BrandDao
	brandMapper *mapper.BrandMapper
}

func NewBrandRepositoryPostgres(db *sql.DB) *BrandRepositoryPostgres {
	return &BrandRepositoryPostgres{
		db:          db,
		brandDao:    dao.NewBrandDao(db),
		brandMapper: mapper.NewBrandMapper(),
	}
}

func (br *BrandRepositoryPostgres) GetAll(ctx context.Context, limit, page int32) ([]model.Brand, error) {
	dtos, err := br.brandDao.GetBrands(ctx, sqlc.GetBrandsParams{
		Limit:  limit,
		Offset: limit * (page - 1),
	})
	if err != nil {
		return nil, err
	}
	entities := make([]model.Brand, 0)
	for _, dto := range dtos {
		entities = append(entities, br.brandMapper.FromDTOToEntity(dto))
	}

	return entities, nil

}

func (br *BrandRepositoryPostgres) Create(ctx context.Context, name string) (uuid.UUID, error) {
	return br.brandDao.CreateBrand(ctx, name)
}

func (br *BrandRepositoryPostgres) GetOne(ctx context.Context, brandID uuid.UUID) (model.Brand, error) {
	dto, err := br.brandDao.GetBrandById(ctx, brandID)
	if err != nil {
		return model.Brand{}, err
	}

	return br.brandMapper.FromDTOToEntity(dto), nil
}

func (br *BrandRepositoryPostgres) Update(ctx context.Context, brandID uuid.UUID, name string) error {
	return br.brandDao.UpdateBrand(ctx, brandID, name)
}

func (br *BrandRepositoryPostgres) Delete(ctx context.Context, brandID uuid.UUID) error {
	return br.brandDao.DeleteBrand(ctx, brandID)
}
