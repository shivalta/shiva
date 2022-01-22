package factory

import (
	"github.com/spf13/viper"
	"shiva/shiva-auth/cmd/http/middlewares"
	"shiva/shiva-auth/configs/driver"
	d_accounts "shiva/shiva-auth/internal/accounts/delivery"
	r_accounts "shiva/shiva-auth/internal/accounts/repository"
	u_accounts "shiva/shiva-auth/internal/accounts/usecase"

	d_class "shiva/shiva-auth/internal/class/delivery"
	r_class "shiva/shiva-auth/internal/class/repository"
	u_class "shiva/shiva-auth/internal/class/usecase"

	d_categories "shiva/shiva-auth/internal/categories/delivery"
	r_categories "shiva/shiva-auth/internal/categories/repository"
	u_categories "shiva/shiva-auth/internal/categories/usecase"
)

type PresenterHTTP struct {
	Accounts   *d_accounts.Http
	Categories *d_categories.Http
	Class      *d_class.Http
}

func InitFactoryHTTP() PresenterHTTP {
	configJWT := middlewares.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	accountsRepo := r_accounts.NewAccountRepo(driver.Psql)
	accountsUsecase := u_accounts.NewAccountUsecase(accountsRepo, &configJWT)
	accountsDelivery := d_accounts.NewAccountsHandler(accountsUsecase)

	classRepo := r_class.NewClassRepo(driver.Psql)
	classUsecase := u_class.NewClassUsecase(classRepo, driver.S3Uploader)
	classDelivery := d_class.NewClassHandler(classUsecase)

	categoriesRepo := r_categories.NewCategoriesRepo(driver.Psql)
	categoriesUsecase := u_categories.NewCategoriesUsecase(categoriesRepo, driver.S3Uploader, classUsecase)
	categoriesDelivery := d_categories.NewCategoriesHandler(categoriesUsecase)

	return PresenterHTTP{
		Accounts:   accountsDelivery,
		Class:      classDelivery,
		Categories: categoriesDelivery,
	}
}
