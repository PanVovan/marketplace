package entity

/*
 * Product model, DB table: "products"
 */
type Product struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Price   int    `json:"price" db:"price"`
	Rating  int    `json:"rating" db:"rating"`
	BrandID int    `json:"brand_id" db:"brand_id"`
}

/*
 * Product image model, DB table: "product_images"
 * One product can have many images
 */
type ProductImage struct {
	ID        int    `json:"-" db:"id"`
	ProductID int    `json:"-" db:"product_id"`
	Path      string `json:"path" db:"path"`
}

/*
 * Binding between the product and the category
 * through the intermediate table "product_categories"
 *
 * This table has composable primary key (category_id + product_id)
 */
type ProductCategory struct {
	CategoryID int `json:"category_id" db:"category_id"`
	ProductID  int `json:"product_id" db:"product_id"`
}

/*
 * Category model, DB table: "categories"
 */
type Category struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

/*
 * Brand model, DB table: "brands"
 */
type Brand struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

/*
 * Binding between the product and the user
 * through the intermediate table "product_category"
 *
 * This table has composable primary key (user_id + product_id)
 */
type Rating struct {
	UserId    int `json:"user_id" db:"user_id"`
	ProductID int `json:"product_id" db:"product_id"`
	Rate      int `json:"rate" db:"rate"`
}

/*
 * Product property model, DB table: "product_properties"
 * One product can have many propetries
 */
type ProductProperty struct {
	ID        int    `json:"id" db:"id"`
	ProductID int    `json:"product_id" db:"product_id"`
	Name      string `json:"name" db:"name"`
	Value     string `json:"value" db:"value"`
}
