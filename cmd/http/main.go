package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	v1.GET("/users", f.Accounts.GetAll, middleware.JWTWithConfig(f.ConfigJWT), middlewares.IsAdmin)
	v1.POST("/users", f.Accounts.Create)
	v1.GET("/users/:userId", f.Accounts.GetById, middleware.JWTWithConfig(f.ConfigJWT), middlewares.IsUserId)
	v1.DELETE("/users/:userId", f.Accounts.Delete, middleware.JWTWithConfig(f.ConfigJWT), middlewares.IsUserId)
	v1.PUT("/users/:userId", f.Accounts.Update, middleware.JWTWithConfig(f.ConfigJWT), middlewares.IsUserId)

	//PRODUCT CLASS ENDPOINT
	v1.GET("/class", f.Class.GetAll)
	v1.POST("/class", f.Class.Create, middleware.JWTWithConfig(f.ConfigJWT), middlewares.IsAdmin)
	v1.GET("/class/:id", f.Class.GetById)
	v1.DELETE("/class/:id", f.Class.Delete, middleware.JWTWithConfig(f.ConfigJWT), middlewares.IsAdmin)
	v1.PUT("/class/:id", f.Class.Update, middleware.JWTWithConfig(f.ConfigJWT), middlewares.IsAdmin)

	//PRODUCT CATEGORY ENDPOINT
	v1.GET("/categories", f.Categories.GetAll)
	v1.POST("/categories", f.Categories.Create, middleware.JWTWithConfig(f.ConfigJWT), middlewares.IsAdmin)
	v1.GET("/categories/:id", f.Categories.GetById)
	v1.DELETE("/categories/:id", f.Categories.Delete, middleware.JWTWithConfig(f.ConfigJWT), middlewares.IsAdmin)
	v1.PUT("/categories/:id", f.Categories.Update, middleware.JWTWithConfig(f.ConfigJWT), middlewares.IsAdmin)

	//PRODUCTS ENDPOINT
	v1.GET("/products", f.Products.GetAll)
	v1.POST("/products", f.Products.Create, middleware.JWTWithConfig(f.ConfigJWT), middlewares.IsAdmin)
	v1.GET("/products/:id", f.Products.GetById)
	v1.DELETE("/products/:id", f.Products.Delete, middleware.JWTWithConfig(f.ConfigJWT), middlewares.IsAdmin)
	v1.PUT("/products/:id", f.Products.Update, middleware.JWTWithConfig(f.ConfigJWT), middlewares.IsAdmin)

	//CHECKOUT ENDPOINT
	v1.POST("/checkout", f.Orders.Checkout)
	v1.POST("/payment", f.Orders.CreateVA, middleware.JWTWithConfig(f.ConfigJWT))

	v1.GET("/payment-list", f.Orders.PaymentMethod)

	err := e.Start(":1111")
	if err != nil {
		return
	}
}
