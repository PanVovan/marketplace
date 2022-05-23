-- name: CreateProductCategory :exec
INSERT INTO products_categories(products_id, categories_id) VALUES ($1,$2);

-- name: GetProductsCategories :many
SELECT products_id, categories_id FROM products_categories AS products_categories LIMIT $1 OFFSET $2;

-- name: GetProductsCategoriesByCategoryID :many
SELECT products_id, categories_id FROM products_categories AS products_categories WHERE categories_id = $1 LIMIT $2 OFFSET $3;

-- name: GetProductsCategoriesByProductID :many
SELECT products_id, categories_id FROM products_categories AS products_categories WHERE products_id = $1 LIMIT $2 OFFSET $3;


-- name: DeleteProductCategoryByProductID :exec
DELETE FROM products_categories WHERE products_id = $1;

-- name: DeleteProductCategoryByCategoryID :exec
DELETE FROM products_categories WHERE categories_id = $1;

-- name: DeleteProductCategory :exec
DELETE FROM products_categories WHERE products_id = $1 AND categories_id = $2;