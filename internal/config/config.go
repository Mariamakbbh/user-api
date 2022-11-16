package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() {

	viper.BindEnv("server.port", "PORT")
	viper.BindEnv("database.user", "DB_USER")
	viper.BindEnv("database.password", "DB_PWD")
	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.port", "DB_PORT")
	viper.BindEnv("database.name", "DB_NAME")

}
