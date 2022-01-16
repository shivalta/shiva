package factory

import (
	"github.com/spf13/viper"
	"shiva/shiva-auth/cmd/http"
	"shiva/shiva-auth/configs/driver"
	d_accounts "shiva/shiva-auth/internal/accounts/delivery"
	r_accounts "shiva/shiva-auth/internal/accounts/repository"
	u_accounts "shiva/shiva-auth/internal/accounts/usecase"
)

type PresenterHTTP struct {
	Accounts *d_accounts.Http
}

func InitFactoryHTTP() PresenterHTTP {
	configJWT := http.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	accountsRepo := r_accounts.NewAccountRepo(driver.Psql)
	accountsUsecase := u_accounts.NewAccountUsecase(accountsRepo, &configJWT)
	accountsDelivery := d_accounts.NewAccountsHandler(accountsUsecase)
	return PresenterHTTP{
		Accounts: accountsDelivery,
	}
}
