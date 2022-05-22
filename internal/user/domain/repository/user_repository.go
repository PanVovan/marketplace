package repository

import (
	"context"
	"makretplace/internal/user/domain/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user model.CreateUserParams) (uuid.UUID, error)
	GetOne(ctx context.Context, userID uuid.UUID) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
	GetAll(ctx context.Context, limit, page int32) ([]model.User, error)
	Update(ctx context.Context, userID uuid.UUID, params model.UpdateUserParams) error
	Delete(ctx context.Context, userID uuid.UUID) error
	GetInfo(ctx context.Context, userID uuid.UUID) (model.UserInfo, error)
	GetInfoAll(ctx context.Context, limit, page int32) ([]model.UserInfo, error)
}
