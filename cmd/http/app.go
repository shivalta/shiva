package http

import (
	"github.com/labstack/echo/v4"
	"shiva/shiva-auth/factory"
)

func InitHttp() {
	f := factory.InitFactoryHTTP()

	e := echo.New()
	initMiddleware(e)

	v1 := e.Group("api/v1/")

	v1.GET("/users", f.Accounts.GetAll)
	v1.POST("/users", f.Accounts.Create)
	err := e.Start(":1111")
	if err != nil {
		return
	}
}