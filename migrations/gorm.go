package migrations

import (
	"shiva/shiva-auth/configs/driver"
	"shiva/shiva-auth/internal/accounts/repository"
)

func AutoMigrate() {
	err := driver.Psql.AutoMigrate(
		&repository.Users{},
	)
	if err != nil {
		return
	}
}
