package factory

import (
	"github.com/spf13/viper"
	"shiva/shiva-auth/cmd/http/middlewares"
	"shiva/shiva-auth/configs/driver"
	d_accounts "shiva/shiva-auth/internal/accounts/delivery"
	r_accounts "shiva/shiva-auth/internal/accounts/repository"
	u_accounts "shiva/shiva-auth/internal/accounts/usecase"
)

type PresenterHTTP struct {
	Accounts *d_accounts.Http
}

func InitFactoryHTTP() PresenterHTTP {
	configJWT := middlewares.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	//just implement this uploader inner usecase for upload to s3 bucket
	//s3 := driver.ConnectAws()
	//uploader := s3manager.NewUploader(s3)

	accountsRepo := r_accounts.NewAccountRepo(driver.Psql)
	accountsUsecase := u_accounts.NewAccountUsecase(accountsRepo, &configJWT)
	accountsDelivery := d_accounts.NewAccountsHandler(accountsUsecase)
	return PresenterHTTP{
		Accounts: accountsDelivery,
	}
}
