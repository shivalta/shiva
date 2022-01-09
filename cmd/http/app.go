package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"shiva/shiva-auth/factory"
)

func InitHttp() {
	f := factory.InitFactoryHTTP()

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/", f.Accounts.GetAll)
	err := e.Start(":1111")
	if err != nil {
		return
	}
}
