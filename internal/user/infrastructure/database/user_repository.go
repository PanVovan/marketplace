package database

import (
	"context"
	"database/sql"
	"makretplace/internal/user/domain/mapper"
	"makretplace/internal/user/domain/model"
	"makretplace/internal/user/infrastructure/database/dao"
	"makretplace/internal/user/infrastructure/database/sqlc"

	"github.com/google/uuid"
)

type UserRepositoryPostgres struct {
	db             *sql.DB
	userDao        *dao.UserDao
	userMapper     *mapper.UserMapper
	UserInfoMapper *mapper.UserInfoMapper
}

func NewUserRepositoryPostgres(db *sql.DB) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{
		db:             db,
		userDao:        dao.NewUserDao(db),
		userMapper:     mapper.NewUserMapper(),
		UserInfoMapper: mapper.NewUserInfoMapper(),
	}
}

func (us *UserRepositoryPostgres) Create(ctx context.Context, user model.CreateUserParams) (uuid.UUID, error) {
	return us.userDao.CreateUser(ctx, sqlc.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
}

func (us *UserRepositoryPostgres) GetOne(ctx context.Context, userID uuid.UUID) (model.User, error) {
	userDto, err := us.userDao.GetUserByID(ctx, userID)
	if err != nil {
		return model.User{}, err
	}
	user := us.userMapper.FromDTOToEntity(userDto)
	return user, nil
}

func (us *UserRepositoryPostgres) GetByEmail(ctx context.Context, email string) (model.User, error) {
	userDto, err := us.userDao.GetUserByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}
	user := us.userMapper.FromDTOToEntity(userDto)
	return user, nil
}

func (us *UserRepositoryPostgres) GetAll(ctx context.Context, limit int32, page int32) ([]model.User, error) {
	usersDto, err := us.userDao.GetUsers(ctx, sqlc.GetUsersParams{
		Limit:  limit,
		Offset: limit * (page - 1),
	})

	if err != nil {
		return nil, err
	}

	users := make([]model.User, 0)

	for _, userDto := range usersDto {
		users = append(users, us.userMapper.FromDTOToEntity(userDto))
	}

	return users, nil
}

func (us *UserRepositoryPostgres) Update(ctx context.Context, userID uuid.UUID, params model.UpdateUserParams) error {
	return us.userDao.UpdateUser(ctx, userID, dao.UpdateUserParams{
		Name:     params.Name,
		Password: params.Password,
		Email:    params.Email,
	})
}

func (us *UserRepositoryPostgres) Delete(ctx context.Context, userID uuid.UUID) error {
	return us.userDao.DeleteUser(ctx, userID)
}

func (us *UserRepositoryPostgres) GetInfo(ctx context.Context, userID uuid.UUID) (model.UserInfo, error) {
	userDto, err := us.userDao.GetUserInfoByID(ctx, userID)
	if err != nil {
		return model.UserInfo{}, err
	}
	user := us.UserInfoMapper.FromDTOToEntity(userDto)
	return user, nil
}

func (us *UserRepositoryPostgres) GetInfoAll(ctx context.Context, limit int32, page int32) ([]model.UserInfo, error) {
	usersDto, err := us.userDao.GetUsersInfo(ctx, dao.GetUsersInfoParams{
		Limit:  limit,
		Offset: limit * (page - 1),
	})

	if err != nil {
		return nil, err
	}

	users := make([]model.UserInfo, 0)

	for _, userDto := range usersDto {
		users = append(users, us.UserInfoMapper.FromDTOToEntity(userDto))
	}

	return users, nil
}
