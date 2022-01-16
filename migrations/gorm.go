package migrations

import (
	"shiva/shiva-auth/configs/driver"
	"shiva/shiva-auth/internal/accounts/repository"
)

func AutoMigrate() {
	err := driver.Psql.AutoMigrate(
		&repository.Users{},
		&repository.Admin{},
	)
	if err != nil {
		return
	}
}
