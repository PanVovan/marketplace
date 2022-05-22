package mapper

import (
	"makretplace/internal/user/domain/model"
	"makretplace/internal/user/infrastructure/database/sqlc"
)

type UserInfoMapper struct{}

func NewUserInfoMapper() *UserInfoMapper {
	return &UserInfoMapper{}
}

func (s *UserInfoMapper) FromDTOToEntity(dto sqlc.UserInfo) model.UserInfo {
	return model.UserInfo{
		ID:   dto.ID,
		Name: dto.Name,
	}
}

func (s *UserInfoMapper) FromEntityToDTO(model model.UserInfo) sqlc.UserInfo {
	return sqlc.UserInfo{
		ID:   model.ID,
		Name: model.Name,
	}
}
