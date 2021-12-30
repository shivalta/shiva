package main

import (
	"github.com/spf13/viper"
	"shiva/shiva-auth/configs/driver"
)

func init() {
	viper.SetConfigFile(`configs/config.yaml`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	driver.SetupDatabasePostgres()
}

func main() {
	//e := echo.New()
	//fmt.Println(viper.GetString("databases.postgres.host"))

}
