package mapper

import (
	"makretplace/internal/basket/domain/aggregate"
	"makretplace/internal/basket/domain/model"
	"makretplace/internal/basket/infrastructure/database/sqlc"
	"makretplace/internal/product/domain/mapper"

	"github.com/google/uuid"
)

type BasketProductMapper struct {
	productInfoMapper *mapper.ProductInfoMapper
	imageMapper       *mapper.ProductImageMapper
}

func NewBasketProductMapper(
	productInfoMapper *mapper.ProductInfoMapper,
	imageMapper *mapper.ProductImageMapper,
) *BasketProductMapper {
	return &BasketProductMapper{
		productInfoMapper: productInfoMapper,
		imageMapper:       imageMapper,
	}
}

func (b *BasketProductMapper) FromDTOToEntity(dto aggregate.BasketProduct) model.BasketProduct {
	product := b.productInfoMapper.FromDTOToEntity(*dto.Product)
	image := b.imageMapper.FromDTOToEntity(*dto.Image)
	return model.BasketProduct{
		ID:       dto.BasketProduct.ID,
		Product:  product,
		Image:    image,
		Quantity: dto.BasketProduct.Quantity,
	}
}

func (b *BasketProductMapper) FromEntityToDTO(model model.BasketProduct, basketID uuid.UUID) aggregate.BasketProduct {

	product := b.productInfoMapper.FromEntityToDTO(model.Product)
	image := b.imageMapper.FromEntityToDTO(model.Image)
	image.ProductID = model.Product.ID

	return aggregate.BasketProduct{
		Product: &product,
		BasketProduct: &sqlc.BasketProduct{
			ID:        model.ID,
			BasketID:  basketID,
			ProductID: model.Product.ID,
			Quantity:  model.Quantity,
		},
		Image: &image,
	}
}
