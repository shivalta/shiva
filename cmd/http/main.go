package http

import (
	"github.com/labstack/echo/v4"
	"shiva/shiva-auth/cmd/http/middlewares"
	"shiva/shiva-auth/factory"
)

func InitHttp() {
	f := factory.InitFactoryHTTP()

	e := echo.New()

	middlewares.InitMiddleware(e)

	v1 := e.Group("api/v1")
	//AUTH ENDPOINT
	v1.POST("/auth/login", f.Accounts.Login)
	v1.POST("/verify", f.Accounts.Verify)

	//USERS ENDPOINT
	v1.GET("/users", f.Accounts.GetAll)
	v1.POST("/users", f.Accounts.Create)
	v1.GET("/users/:userId", f.Accounts.GetById)
	v1.DELETE("/users/:userId", f.Accounts.Delete)
	v1.PUT("/users/:userId", f.Accounts.Update)

	//PRODUCT CLASS ENDPOINT
	v1.GET("/class", f.Class.GetAll)
	v1.POST("/class", f.Class.Create)
	v1.GET("/class/:id", f.Class.GetById)
	v1.DELETE("/class/:id", f.Class.Delete)
	v1.PUT("/class/:id", f.Class.Update)

	//PRODUCT CATEGORY ENDPOINT
	v1.GET("/categories", f.Categories.GetAll)
	v1.POST("/categories", f.Categories.Create)
	v1.GET("/categories/:id", f.Categories.GetById)
	v1.DELETE("/categories/:id", f.Categories.Delete)
	v1.PUT("/categories/:id", f.Categories.Update)

	//PRODUCTS ENDPOINT
	v1.GET("/products", f.Products.GetAll)
	v1.POST("/products", f.Products.Create)
	v1.GET("/products/:id", f.Products.GetById)
	v1.DELETE("/products/:id", f.Products.Delete)
	v1.PUT("/products/:id", f.Products.Update)

	//CHECKOUT ENDPOINT
	v1.POST("/checkout", f.Orders.Checkout)

	v1.GET("/payment-list", f.Orders.PaymentMethod)

	err := e.Start(":1111")
	if err != nil {
		return
	}
}
