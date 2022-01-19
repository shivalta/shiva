package driver

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var Psql *gorm.DB

func SetupDatabasePostgres() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		viper.GetString(`databases.postgres.host`),
		viper.GetString(`databases.postgres.user`),
		viper.GetString(`databases.postgres.password`),
		viper.GetString(`databases.postgres.dbname`),
		viper.GetString(`databases.postgres.port`),
		viper.GetString(`databases.postgres.sslmode`),
		viper.GetString(`databases.postgres.timezone`),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Logger = logger.Default.LogMode(logger.Info)
	if err != nil {
		panic(err.Error)
	}
	log.Println("Postgres Connected!")
	Psql = db
}
