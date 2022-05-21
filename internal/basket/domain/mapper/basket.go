package mapper

import (
	"makretplace/internal/basket/domain/aggregate"
	"makretplace/internal/basket/domain/model"
	"makretplace/internal/basket/infrastructure/database/sqlc"
	"makretplace/internal/user/domain/mapper"

	"github.com/google/uuid"
)

type BasketMapper struct {
	basketProductMapper *BasketProductMapper
	userInfoMapper      *mapper.UserInfoMapper
}

func NewBasketMapper(
	basketProductMapper *BasketProductMapper,
	userInfoMapper *mapper.UserInfoMapper,
) *BasketMapper {
	return &BasketMapper{
		basketProductMapper: basketProductMapper,
		userInfoMapper:      userInfoMapper,
	}
}

func (b *BasketMapper) FromDTOToEntity(dto aggregate.Basket) model.Basket {

	var products []model.BasketProduct
	for _, product := range dto.Products {
		products = append(products, b.basketProductMapper.FromDTOToEntity(*product))
	}

	return model.Basket{
		ID:             dto.Basket.ID,
		User:           b.userInfoMapper.FromDTOToEntity(*dto.User),
		BasketProducts: products,
	}
}

func (b *BasketMapper) FromEntityToDTO(model model.Basket) aggregate.Basket {

	user := b.userInfoMapper.FromEntityToDTO(model.User)

	var products []*aggregate.BasketProduct

	for _, product := range model.BasketProducts {
		pr := b.basketProductMapper.FromEntityToDTO(product, model.User.ID)
		products = append(products, &pr)
	}

	return aggregate.Basket{
		Basket: &sqlc.Basket{
			ID:     model.ID,
			UserID: uuid.NullUUID{UUID: model.User.ID},
		},
		User:     &user,
		Products: products,
	}
}
