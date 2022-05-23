package database

import (
	"context"
	"database/sql"
	brand_mapper "makretplace/internal/brand/domain/mapper"
	brand_dao "makretplace/internal/brand/infrastructure/database/dao"
	brand_dto "makretplace/internal/brand/infrastructure/database/sqlc"
	category_mapper "makretplace/internal/category/domain/mapper"
	category_dao "makretplace/internal/category/infrastructure/database/dao"
	"makretplace/internal/product/domain/aggregate"
	product_mapper "makretplace/internal/product/domain/mapper"
	"makretplace/internal/product/domain/model"
	product_repo "makretplace/internal/product/domain/repository"
	"makretplace/internal/product/infrastructure/database/dao"
	"makretplace/internal/product/infrastructure/database/sqlc"
	seller_mapper "makretplace/internal/seller/domain/mapper"
	seller_dao "makretplace/internal/seller/infrastructure/database/dao"

	"github.com/google/uuid"
)

type ProductRepositoryPostgres struct {
	db                 *sql.DB
	dao                *dao.ProductDao
	propertyRepository *product_repo.ProductPropertyRepository
	sellerDao          *seller_dao.SellerDao
	imageRepository    *product_repo.ProductImageRepository
	brandDao           *brand_dao.BrandDao
	categoryDao        *category_dao.CategoryDao
	productMapper      *product_mapper.ProductMapper
	productInfoMapper  *product_mapper.ProductInfoMapper
}

func NewProductRepositoryPostgres(
	db *sql.DB,
	propertyRepository *product_repo.ProductPropertyRepository,
	imageRepository *product_repo.ProductImageRepository,
) *ProductRepositoryPostgres {
	return &ProductRepositoryPostgres{
		db:                 db,
		propertyRepository: propertyRepository,
		sellerDao:          seller_dao.NewSellerDao(db),
		brandDao:           brand_dao.NewBrandDao(db),
		categoryDao:        category_dao.NewCategoryDao(db),
		imageRepository:    imageRepository,
		dao:                dao.NewProductDao(db),
		productMapper: product_mapper.NewProductMapper(
			brand_mapper.NewBrandMapper(),
			category_mapper.NewCategoryMapper(),
			product_mapper.NewProductPropertyMapper(),
			product_mapper.NewProductImageMapper(),
			seller_mapper.NewSellerInfoMapper(),
		),
		productInfoMapper: product_mapper.NewProductInfoMapper(
			brand_mapper.NewBrandMapper(),
			category_mapper.NewCategoryMapper(),
			seller_mapper.NewSellerInfoMapper(),
		),
	}
}

func (pr *ProductRepositoryPostgres) Create(ctx context.Context, product model.CreateProductParams) (uuid.UUID, error) {

	brandID := uuid.NullUUID{
		UUID:  uuid.Nil,
		Valid: false,
	}
	if product.BrandID != nil {
		brandID.UUID = *product.BrandID
		brandID.Valid = true
	}

	productId, err := pr.dao.CreateProduct(ctx, sqlc.CreateProductParams{
		Name:     product.Name,
		Price:    product.Price,
		Rating:   product.Rating,
		BrandID:  brandID,
		SellerID: product.SellerID,
		Amount:   product.Amount,
	})

	if err != nil {
		return uuid.Nil, err
	}

	//TODO: Create file uploading
	if product.Images != nil {
		for _, image := range product.Images {
			_, err := (*pr.imageRepository).Create(ctx, productId, image.File)
			if err != nil {
				return uuid.Nil, err
			}
		}
	}

	if product.Properties != nil {
		for _, property := range product.Properties {

			_, err := (*pr.propertyRepository).Create(ctx, productId, model.CreateProductPropertyParams{
				Name:  property.Name,
				Value: property.Value,
			})
			if err != nil {
				return uuid.Nil, err
			}
		}
	}

	if product.Categories != nil {
		for _, category := range product.Categories {
			err := pr.dao.CreateProductCategory(ctx, sqlc.CreateProductCategoryParams{
				ProductsID:   productId,
				CategoriesID: category,
			})
			if err != nil {
				return uuid.Nil, err
			}

		}
	}

	return productId, nil
}

func (pr *ProductRepositoryPostgres) GetOne(ctx context.Context, productID uuid.UUID) (model.Product, error) {
	prod, err := pr.dao.GetProductByID(ctx, productID)
	if err != nil {
		return model.Product{}, err
	}

	images, err := pr.dao.GetImagesByProductID(ctx, productID)
	if err != nil {
		return model.Product{}, err
	}

	imag_ptr := make([]*sqlc.ProductImage, 0)
	for _, ptr := range images {
		imag_ptr = append(imag_ptr, &ptr)
	}

	properties, err := pr.dao.GetProductPropertiesByProductID(ctx, productID)
	if err != nil {
		return model.Product{}, err
	}

	prop_ptr := make([]*sqlc.ProductProperty, 0)
	for _, ptr := range properties {
		prop_ptr = append(prop_ptr, &ptr)
	}

	ctg, err := pr.dao.GetCategoriesByProductID(ctx, productID)
	if err != nil {
		return model.Product{}, err
	}

	var brand brand_dto.Brand

	if prod.BrandID.Valid {
		brand, err = pr.brandDao.GetBrandById(ctx, prod.BrandID.UUID)
		if err != nil {
			return model.Product{}, err
		}
	}

	seller, err := pr.sellerDao.GetSellerInfoByID(ctx, prod.SellerID)
	if err != nil {
		return model.Product{}, err
	}

	product := aggregate.Product{
		Product:    &prod,
		Images:     imag_ptr,
		Brand:      &brand,
		Categories: ctg,
		SellerInfo: &seller,
		Properties: prop_ptr,
	}

	return pr.productMapper.FromDTOToEntity(product), nil
}

func (pr *ProductRepositoryPostgres) GetAll(ctx context.Context, specs product_repo.GetProductsQuerySpecs, limit, page int32) ([]model.ProductInfo, error) {
	products, err := pr.dao.GetProducts(ctx, dao.GetProductsParams{
		Limit:  limit,
		Offset: limit * (page - 1),
	}, dao.GetProductsQuerySpecs{
		BrandID:      specs.BrandID,
		CategoriesID: specs.CategoriesID,
		SellerID:     specs.SellerID,
	})

	if err != nil {
		return nil, err
	}

	models := make([]model.ProductInfo, 0)
	for _, prod := range products {
		productID := prod.ID

		ctg, err := pr.dao.GetCategoriesByProductID(ctx, productID)
		if err != nil {
			return nil, err
		}

		var brand brand_dto.Brand

		if prod.BrandID.Valid {
			brand, err = pr.brandDao.GetBrandById(ctx, prod.BrandID.UUID)
			if err != nil {
				return nil, err
			}
		}

		seller, err := pr.sellerDao.GetSellerInfoByID(ctx, prod.SellerID)
		if err != nil {
			return nil, err
		}

		product := aggregate.ProductInfo{
			Product:    &prod,
			Brand:      &brand,
			Categories: ctg,
			SellerInfo: &seller,
		}

		models = append(models, pr.productInfoMapper.FromDTOToEntity(product))

	}
	return models, nil
}

func (pr *ProductRepositoryPostgres) Update(ctx context.Context, productID uuid.UUID, params model.UpdateProductParams) error {
	err := pr.dao.UpdateProduct(ctx, dao.UpdateProductParams{
		Name:     params.Name,
		Price:    params.Price,
		Rating:   params.Rating,
		BrandID:  params.BrandID,
		SellerID: params.SellerID,
		Amount:   params.Amount,
	}, productID)
	if err != nil {
		return err
	}

	if params.Categories != nil {
		pr.dao.DeleteProductCategoryByProductID(ctx, productID)
		for _, category := range params.Categories {
			err := pr.dao.CreateProductCategory(ctx, sqlc.CreateProductCategoryParams{
				ProductsID:   productID,
				CategoriesID: category,
			})
			if err != nil {
				return err
			}
		}
	}

	if params.Properties != nil {
		pr.dao.DeleteProductPropertiesByProductID(ctx, productID)
		for _, property := range params.Properties {
			_, err := (*pr.propertyRepository).Create(ctx, productID, model.CreateProductPropertyParams{
				Name:  property.Name,
				Value: property.Value,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
func (pr *ProductRepositoryPostgres) Delete(ctx context.Context, productID uuid.UUID) error {
	pr.dao.DeleteProductPropertiesByProductID(ctx, productID)
	pr.dao.DeleteProductCategoryByProductID(ctx, productID)
	return pr.dao.DeleteProduct(ctx, productID)
}
