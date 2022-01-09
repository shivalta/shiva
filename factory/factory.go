package factory

import (
	"shiva/shiva-auth/configs/driver"
	d_accounts "shiva/shiva-auth/internal/accounts/delivery"
	r_accounts "shiva/shiva-auth/internal/accounts/repository"
	u_accounts "shiva/shiva-auth/internal/accounts/usecase"
)

type PresenterHTTP struct {
	Accounts *d_accounts.Http
}

func InitFactoryHTTP() PresenterHTTP {
	accountsRepo := r_accounts.NewAccountRepo(driver.Psql)
	accountsUsecase := u_accounts.NewAccountUsecase(accountsRepo)
	accountsDelivery := d_accounts.NewAccountsHandler(accountsUsecase)
	return PresenterHTTP{
		Accounts: accountsDelivery,
	}
}
