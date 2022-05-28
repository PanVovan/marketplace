package application

import (
	"context"
	"makretplace/internal/product/domain/model"
	"makretplace/internal/product/domain/repository"

	"github.com/google/uuid"
)

type GetAllUseCase struct {
	repository repository.ProductRepository
}

func NewGetAllUseCase(
	repository repository.ProductRepository,
) *GetAllUseCase {
	return &GetAllUseCase{
		repository: repository,
	}
}

func (ga *GetAllUseCase) Execute(ctx context.Context, sellerID, brandId *string, categoriesId []*string, limit, page int32) ([]model.ProductInfo, error) {
	specs := repository.GetProductsQuerySpecs{
		SellerID:     nil,
		CategoriesID: nil,
		BrandID:      nil,
	}

	if brandId != nil {
		brand, err := uuid.Parse(*brandId)
		if err != nil {
			return nil, err
		}
		specs.BrandID = &brand
	}
	if categoriesId == nil {
		categories := make([]*uuid.UUID, 0)
		for _, categoryId := range categoriesId {
			category, err := uuid.Parse(*categoryId)
			if err != nil {
				return nil, err
			}
			categories = append(categories, &category)

		}
		specs.CategoriesID = categories
	}

	if sellerID != nil {
		seller, err := uuid.Parse(*sellerID)
		if err != nil {
			return nil, err
		}
		specs.SellerID = &seller
	}

	return ga.repository.GetAll(ctx, specs, limit, page)
}
