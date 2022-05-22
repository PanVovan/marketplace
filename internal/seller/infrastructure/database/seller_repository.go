package database

import (
	"context"
	"database/sql"
	"makretplace/internal/seller/domain/mapper"
	"makretplace/internal/seller/domain/model"
	"makretplace/internal/seller/infrastructure/database/dao"
	"makretplace/internal/seller/infrastructure/database/sqlc"

	"github.com/google/uuid"
)

type SellerRepositoryPostgres struct {
	db               *sql.DB
	sellerDao        *dao.SellerDao
	sellerMapper     *mapper.SellerMapper
	sellerInfoMapper *mapper.SellerInfoMapper
}

func NewSellerRepositoryPostgres(db *sql.DB) *SellerRepositoryPostgres {
	return &SellerRepositoryPostgres{
		db:               db,
		sellerDao:        dao.NewSellerDao(db),
		sellerMapper:     mapper.NewSellerMapper(),
		sellerInfoMapper: mapper.NewSellerInfoMapper(),
	}
}

func (us *SellerRepositoryPostgres) Create(ctx context.Context, user model.CreateSellerParams) (uuid.UUID, error) {
	return us.sellerDao.CreateSeller(ctx, sqlc.CreateSellerParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
}

func (us *SellerRepositoryPostgres) GetOne(ctx context.Context, userID uuid.UUID) (model.Seller, error) {
	userDto, err := us.sellerDao.GetSellerByID(ctx, userID)
	if err != nil {
		return model.Seller{}, err
	}
	user := us.sellerMapper.FromDTOToEntity(userDto)
	return user, nil
}

func (us *SellerRepositoryPostgres) GetByEmail(ctx context.Context, email string) (model.Seller, error) {
	sellerDto, err := us.sellerDao.GetSellerByEmail(ctx, email)
	if err != nil {
		return model.Seller{}, err
	}
	user := us.sellerMapper.FromDTOToEntity(sellerDto)
	return user, nil
}

func (us *SellerRepositoryPostgres) GetAll(ctx context.Context, limit int32, page int32) ([]model.Seller, error) {
	sellersDto, err := us.sellerDao.GetSellers(ctx, sqlc.GetSellersParams{
		Limit:  limit,
		Offset: limit * (page - 1),
	})

	if err != nil {
		return nil, err
	}

	sellers := make([]model.Seller, 0)

	for _, sellerDto := range sellersDto {
		sellers = append(sellers, us.sellerMapper.FromDTOToEntity(sellerDto))
	}

	return sellers, nil
}

func (us *SellerRepositoryPostgres) Update(ctx context.Context, sellerID uuid.UUID, params model.UpdateSellerParams) error {
	return us.sellerDao.UpdateSeller(ctx, sellerID, dao.UpdateSellerParams{
		Name:     params.Name,
		Password: params.Password,
		Email:    params.Email,
	})
}

func (us *SellerRepositoryPostgres) Delete(ctx context.Context, sellerID uuid.UUID) error {
	return us.sellerDao.DeleteSeller(ctx, sellerID)
}

func (us *SellerRepositoryPostgres) GetInfo(ctx context.Context, sellerID uuid.UUID) (model.SellerInfo, error) {
	sellerDto, err := us.sellerDao.GetSellerInfoByID(ctx, sellerID)
	if err != nil {
		return model.SellerInfo{}, err
	}
	seller := us.sellerInfoMapper.FromDTOToEntity(sellerDto)
	return seller, nil
}

func (us *SellerRepositoryPostgres) GetInfoAll(ctx context.Context, limit int32, page int32) ([]model.SellerInfo, error) {
	sellersDto, err := us.sellerDao.GetSellersInfo(ctx, dao.GetSellersInfoParams{
		Limit:  limit,
		Offset: limit * (page - 1),
	})

	if err != nil {
		return nil, err
	}

	sellers := make([]model.SellerInfo, 0)

	for _, sellerDto := range sellersDto {
		sellers = append(sellers, us.sellerInfoMapper.FromDTOToEntity(sellerDto))
	}

	return sellers, nil
}
