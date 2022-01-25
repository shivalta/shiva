package factory

import (
	"github.com/labstack/echo/v4/middleware"
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

	d_products "shiva/shiva-auth/internal/products/delivery"
	r_products "shiva/shiva-auth/internal/products/repository"
	u_products "shiva/shiva-auth/internal/products/usecase"

	d_orders "shiva/shiva-auth/internal/orders/delivery"
	r_orders "shiva/shiva-auth/internal/orders/repository"
	u_orders "shiva/shiva-auth/internal/orders/usecase"
)

type PresenterHTTP struct {
	Accounts   *d_accounts.Http
	Categories *d_categories.Http
	Class      *d_class.Http
	Products   *d_products.Http
	Orders     *d_orders.Http
	ConfigJWT  middleware.JWTConfig
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

	productsRepo := r_products.NewProductsRepo(driver.Psql)
	productsUsecase := u_products.NewProductsUsecase(productsRepo, classUsecase, categoriesUsecase)
	productsDelivery := d_products.NewProductsHandler(productsUsecase)

	ordersMockapi := r_orders.NewMockupApi(viper.GetString(`mockapi.base_url`))
	ordersXendit := r_orders.NewXenditAPI(viper.GetString(`xendit.base_url`), viper.GetString(`xendit.api_key`))
	ordersRepo := r_orders.NewOrdersRepo(driver.Psql)
	ordersUsecase := u_orders.NewOrdersUsecase(ordersRepo, ordersXendit, ordersMockapi, productsUsecase)
	ordersDelivery := d_orders.NewOrdersHandler(ordersUsecase)

	return PresenterHTTP{
		Accounts:   accountsDelivery,
		Categories: categoriesDelivery,
		Class:      classDelivery,
		Products:   productsDelivery,
		Orders:     ordersDelivery,
		ConfigJWT:  configJWT.Init(),
	}
}
