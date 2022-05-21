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

type ProductMapper struct {
	brandMapper      *brand_mapper.BrandMapper
	categoryMapper   *category_mapper.CategoryMapper
	propertyMapper   *productPropertyMapper
	imageMapper      *ProductImageMapper
	sellerInfoMapper *seller_mapper.SellerInfoMapper
}

func NewProductMapper(
	brandMapper *brand_mapper.BrandMapper,
	categoryMapper *category_mapper.CategoryMapper,
	propertyMapper *productPropertyMapper,
	imageMapper *ProductImageMapper,
	sellerInfoMapper *seller_mapper.SellerInfoMapper,
) *ProductMapper {
	return &ProductMapper{
		brandMapper:      brandMapper,
		categoryMapper:   categoryMapper,
		propertyMapper:   propertyMapper,
		imageMapper:      imageMapper,
		sellerInfoMapper: sellerInfoMapper,
	}
}

func (p *ProductMapper) FromDTOToEntity(aggregate aggregate.Product) product_model.Product {

	var categories []category_model.Category
	for _, categoryDto := range aggregate.Categories {
		categories = append(categories, p.categoryMapper.FromDTOToEntity(*categoryDto))
	}

	var properties []product_model.ProductProperty
	for _, propertyDto := range aggregate.Properties {
		properties = append(properties, p.propertyMapper.FromDTOToEntity(*propertyDto))
	}

	var images []product_model.ProductImage
	for _, imageDto := range aggregate.Images {
		images = append(images, p.imageMapper.FromDTOToEntity(*imageDto))
	}

	return product_model.Product{
		ID:         aggregate.Product.ID,
		Name:       aggregate.Product.Name,
		Price:      aggregate.Product.Price,
		Rating:     aggregate.Product.Rating,
		Seller:     p.sellerInfoMapper.FromDTOToEntity(*aggregate.SellerInfo),
		Amount:     aggregate.Product.Amount,
		Brand:      p.brandMapper.FromDTOToEntity(*aggregate.Brand),
		Categories: categories,
		Properties: properties,
		Images:     images,
	}
}

func (p *ProductMapper) FromEntityToDTO(model product_model.Product) aggregate.Product {

	brand := p.brandMapper.FromEntityToDTO(brand_model.Brand{
		ID:   model.Brand.ID,
		Name: model.Brand.Name,
	})

	var properties []*product_dto.ProductProperty
	for _, prop := range model.Properties {
		property := p.propertyMapper.FromEntityToDTO(prop)
		property.ProductID = model.ID
		properties = append(properties, &property)
	}

	var categories []*category_dto.Category
	for _, category := range model.Categories {
		c := p.categoryMapper.FromEntityToDTO(category)
		categories = append(categories, &c)
	}

	var images []*product_dto.ProductImage
	for _, image := range model.Images {
		i := p.imageMapper.FromEntityToDTO(image)
		i.ProductID = model.ID
		images = append(images, &i)

	}

	seller := p.sellerInfoMapper.FromEntityToDTO(model.Seller)

	return aggregate.Product{
		Product:    p.fromEntityToProductDTO(model),
		Brand:      &brand,
		Properties: properties,
		Categories: categories,
		Images:     images,
		SellerInfo: &seller,
	}
}

func (p *ProductMapper) fromEntityToProductDTO(model product_model.Product) *product_dto.Product {
	return &product_dto.Product{
		ID:       model.ID,
		Name:     model.Name,
		Price:    model.Price,
		Rating:   model.Rating,
		BrandID:  uuid.NullUUID{UUID: model.Brand.ID},
		SellerID: model.Seller.ID,
		Amount:   model.Amount,
	}
}
