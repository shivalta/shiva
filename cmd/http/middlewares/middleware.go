package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitMiddleware(e *echo.Echo) *echo.Echo {
	e.Pre(middleware.RemoveTrailingSlash())
	return e
}
