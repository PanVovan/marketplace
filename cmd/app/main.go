package main

import (
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

}

func InitConfig() error {
	viper.AddConfigPath("../../config")
	viper.SetConfigFile("config")
	return viper.ReadInConfig()
}
