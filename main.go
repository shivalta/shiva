package main

import (
	"github.com/spf13/viper"
	"shiva/shiva-auth/cmd/http"
	"shiva/shiva-auth/configs/driver"
	"shiva/shiva-auth/migrations"
)

func init() {
	viper.SetConfigFile(`configs/env/config.yaml`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	driver.SetupDatabasePostgres()
	migrations.AutoMigrate()
}

func main() {
	http.InitHttp()
}
