package mapper

import (
	brand_mapper "makretplace/internal/brand/domain/mapper"
	brand_model "makretplace/internal/brand/domain/model"
	category_mapper "makretplace/internal/category/domain/mapper"
	category_model "makretplace/internal/category/domain/model"
	category_dto "makretplace/internal/category/infrastructure/database/sqlc"
	"makretplace/internal/product/domain/aggregate"
	product_model "makretplace/internal/product/domain/model"
	product_dto "makretplace/internal/product/infrastructure/database/sqlc"
	seller_mapper "makretplace/internal/seller/domain/mapper"

	"github.com/google/uuid"
)

type ProductInfoMapper struct {
	brandMapper    *brand_mapper.BrandMapper
	categoryMapper *category_mapper.CategoryMapper

	sellerInfoMapper *seller_mapper.SellerInfoMapper
}

func NewProductInfoMapper(
	brandMapper *brand_mapper.BrandMapper,
	categoryMapper *category_mapper.CategoryMapper,

	sellerInfoMapper *seller_mapper.SellerInfoMapper,
) *ProductInfoMapper {
	return &ProductInfoMapper{
		brandMapper:      brandMapper,
		categoryMapper:   categoryMapper,
		sellerInfoMapper: sellerInfoMapper,
	}
}

func (p *ProductInfoMapper) FromDTOToEntity(aggregate aggregate.ProductInfo) product_model.ProductInfo {

	var categories []category_model.Category
	for _, categoryDto := range aggregate.Categories {
		categories = append(categories, p.categoryMapper.FromDTOToEntity(*categoryDto))
	}

	return product_model.ProductInfo{
		ID:          aggregate.Product.ID,
		Name:        aggregate.Product.Name,
		Price:       aggregate.Product.Price,
		Rating:      aggregate.Product.Rating,
		Description: aggregate.Product.Description,
		Seller:      p.sellerInfoMapper.FromDTOToEntity(*aggregate.SellerInfo),
		Amount:      aggregate.Product.Amount,
		Brand:       p.brandMapper.FromDTOToEntity(*aggregate.Brand),
		Categories:  categories,
	}
}

func (p *ProductInfoMapper) FromEntityToDTO(model product_model.ProductInfo) aggregate.ProductInfo {

	brand := p.brandMapper.FromEntityToDTO(brand_model.Brand{
		ID:   model.Brand.ID,
		Name: model.Brand.Name,
	})

	var categories []*category_dto.Category
	for _, category := range model.Categories {
		c := p.categoryMapper.FromEntityToDTO(category)
		categories = append(categories, &c)
	}

	seller := p.sellerInfoMapper.FromEntityToDTO(model.Seller)

	return aggregate.ProductInfo{
		Product:    p.fromEntityToProductDTO(model),
		Brand:      &brand,
		Categories: categories,
		SellerInfo: &seller,
	}
}

func (p *ProductInfoMapper) fromEntityToProductDTO(model product_model.ProductInfo) *product_dto.Product {
	return &product_dto.Product{
		ID:          model.ID,
		Name:        model.Name,
		Price:       model.Price,
		Rating:      model.Rating,
		Description: model.Description,
		BrandID:     uuid.NullUUID{UUID: model.Brand.ID},
		SellerID:    model.Seller.ID,
		Amount:      model.Amount,
	}
}
