package repository

import (
	"context"
	"makretplace/internal/seller/domain/model"

	"github.com/google/uuid"
)

type SellerRepository interface {
	Create(ctx context.Context, user model.CreateSellerParams) (uuid.UUID, error)
	GetOne(ctx context.Context, userID uuid.UUID) (model.Seller, error)
	GetByEmail(ctx context.Context, email string) (model.Seller, error)
	GetAll(ctx context.Context, limit, page int32) ([]model.Seller, error)
	Update(ctx context.Context, userID uuid.UUID, params model.UpdateSellerParams) error
	Delete(ctx context.Context, userID uuid.UUID) error
	GetInfo(ctx context.Context, userID uuid.UUID) (model.SellerInfo, error)
	GetInfoAll(ctx context.Context, limit, page int32) ([]model.SellerInfo, error)
}
