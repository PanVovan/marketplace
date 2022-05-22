package main

import (
	"makretplace/internal/basket"
	"makretplace/internal/brand"
	"makretplace/internal/category"
	"makretplace/internal/order"
	"makretplace/internal/product"
	"makretplace/internal/rating"
	"makretplace/internal/seller"
	"makretplace/internal/user"
	marketplace "makretplace/pkg/httpserver"
	"makretplace/pkg/postgres"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	godotenv.Load()
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	routes := mux.NewRouter()

	userModule := &user.Module{}
	userModule.Configure(db, routes)

	sellerModule := &seller.Module{}
	sellerModule.Configure(db, routes)

	productModule := &product.Module{}
	productModule.Configure(db, routes)

	categoryModule := &category.Module{}
	categoryModule.Configure(db, routes)

	brandModule := &brand.Module{}
	brandModule.Configure(db, routes)

	basketModule := &basket.Module{}
	basketModule.Configure(db, routes)

	orderModule := &order.Module{}
	orderModule.Configure(db, routes)

	ratingModule := &rating.Module{}
	ratingModule.Configure(db, routes)

	server := new(marketplace.Server)
	server.Run("localhost", "8080", routes)
}

func InitConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
