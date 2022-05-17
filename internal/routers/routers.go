package v1

import (
	controllers "marketplace/internal/controller/http/v1"

	"github.com/gin-gonic/gin"
)

type Routes struct{}

//TODO: init routes
func (r *Routes) InitRoutes() *gin.Engine {
	router := gin.New()

	InitProductRouters(router.Group("/product"))

	InitCategoryRouters(router.Group("/category"))

	InitBrandRouter(router.Group("/brand"))

	InitUsersRouter(router.Group("/user"))

	return router
}

func InitUsersRouter(user *gin.RouterGroup) {
	userController := new(controllers.User)

	user.GET("/check", userController.Check)
	user.POST("/login", userController.Login)
	user.POST("/signup", userController.SignUp)

	user.GET("/getall", userController.GetAll)
	user.GET("/getone/:id", userController.GetOne)
	user.POST("/create", userController.Create)
	user.PUT("/update/:id", userController.Update)
	user.DELETE("/delete/:id", userController.Delete)
}

func InitBrandRouter(brand *gin.RouterGroup) {
	brandController := new(controllers.Brand)

	brand.GET("/getall", brandController.GetAll)
	brand.GET("/getone/:id", brandController.GetOne)
	brand.POST("/create", brandController.Create)
	brand.PUT("/update/:id", brandController.Update)
	brand.DELETE("/delete/:id", brandController.Delete)
}

func InitCategoryRouters(category *gin.RouterGroup) {
	categoryController := new(controllers.Category)

	category.GET("/getall", categoryController.GetAll)
	category.GET("/getone/:id", categoryController.GetOne)
	category.POST("/create", categoryController.Create)
	category.PUT("/update/:id", categoryController.Update)
	category.DELETE("/delete/:id", categoryController.Delete)
}

func InitProductRouters(product *gin.RouterGroup) {
	productController := new(controllers.Product)

	product.GET("/getall", productController.GetAll)
	product.GET("/getone/:id", productController.GetOne)
	product.POST("/create", productController.Create)
	product.PUT("/update/:id", productController.Update)
	product.DELETE("/delete/:id", productController.Delete)
}
