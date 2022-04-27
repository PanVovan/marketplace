package v1

import (
	"github.com/gin-gonic/gin"
)

type Routes struct {
	Product
	Category
	Brand
}

//TODO: init routes
func (r *Routes) InitRoutes() *gin.Engine {
	router := gin.New()

	product := router.Group("/product")
	{
		productController := new(Product)

		product.GET("/getall", productController.GetAll)
		product.GET("/getone/:id", productController.GetOne)
		product.POST("/create", productController.Create)
		product.PUT("/update/:id", productController.Update)
		product.DELETE("/delete/:id", productController.Delete)
	}

	category := router.Group("/category")
	{
		categoryController := new(Category)

		category.GET("/getall", categoryController.GetAll)
		category.GET("/getone/:id", categoryController.GetOne)
		category.POST("/create", categoryController.Create)
		category.PUT("/update/:id", categoryController.Update)
		category.DELETE("/delete/:id", categoryController.Delete)
	}

	brand := router.Group("/brand")
	{
		brandController := new(Brand)

		brand.GET("/getall", brandController.GetAll)
		brand.GET("/getone/:id", brandController.GetOne)
		brand.POST("/create", brandController.Create)
		brand.PUT("/update/:id", brandController.Update)
		brand.DELETE("/delete/:id", brandController.Delete)
	}

	user := router.Group("/user")
	{
		userController := new(User)

		user.GET("/check", userController.Check)
		user.POST("/login", userController.Login)
		user.POST("/signup", userController.SignUp)

		user.GET("/getall", userController.GetAll)
		user.GET("/getone/:id", userController.GetOne)
		user.POST("/create", userController.Create)
		user.PUT("/update/:id", userController.Update)
		user.DELETE("/delete/:id", userController.Delete)
	}

	return router
}
