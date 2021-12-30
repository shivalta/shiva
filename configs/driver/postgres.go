package driver

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func SetupDatabasePostgres() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		viper.GetString(`databases.postgres.host`),
		viper.GetString(`databases.postgres.user`),
		viper.GetString(`databases.postgres.password`),
		viper.GetString(`databases.postgres.dbname`),
		viper.GetString(`databases.postgres.port`),
		viper.GetString(`databases.postgres.sslmode`),
		viper.GetString(`databases.postgres.timezone`),
	)
	fmt.Println(viper.Get("databases.postgres.host"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Logger = logger.Default.LogMode(logger.Info)
	if err != nil {
		panic(err.Error)
	}
	log.Println("Postgres Connected!")
	return db
}
