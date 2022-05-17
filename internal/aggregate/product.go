package aggregate

import (
	"errors"
	"marketplace/internal/entity"

	"github.com/google/uuid"
)

type Product struct {
	product       *entity.Product
	brand         *entity.Brand
	productImages []*entity.ProductImage
	categories    []*entity.Category
	properties    []*entity.ProductProperty
}

func NewProduct(name string) (Product, error) {
	if name == "" {
		return Product{}, errors.New("Product has to have an valid")
	}
	product := &entity.Product{
		ID:   uuid.Nil,
		Name: name,
	}
	return Product{
		product:       product,
		brand:         nil,
		productImages: make([]*entity.ProductImage, 0),
		categories:    make([]*entity.Category, 0),
		properties:    make([]*entity.ProductProperty, 0),
	}, nil
}

func (p *Product) GetID() uuid.UUID {
	return p.product.ID
}

func (p *Product) GetName() string {
	return p.product.Name
}

func (p *Product) GetPrice() float64 {
	return p.product.Price
}

func (p *Product) GetRating() int {
	return p.product.Rating
}

func (p *Product) GetBrand() *entity.Brand {
	return p.brand
}

func (p *Product) GetProductImages() []*entity.ProductImage {
	return p.productImages
}

func (p *Product) GetCategories() []*entity.Category {
	return p.categories
}

func (p *Product) GetProperties() []*entity.ProductProperty {
	return p.properties
}

func (p *Product) SetID(id uuid.UUID) {
	p.product.ID = id
}

func (p *Product) SetName(name string) {
	p.product.Name = name
}

func (p *Product) SetPrice(price float64) {
	p.product.Price = price
}

func (p *Product) SetRating(rating int) {
	p.product.Rating = rating
}

func (p *Product) SetBrand(brand *entity.Brand) {
	p.brand = brand
}

func (p *Product) SetProductImages(productImages []*entity.ProductImage) {
	p.productImages = productImages
}

func (p *Product) SetCategories(categories []*entity.Category) {
	p.categories = categories
}

func (p *Product) SetProperties(properties []*entity.ProductProperty) {
	p.properties = properties
}
