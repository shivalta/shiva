package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initMiddleware(e *echo.Echo) *echo.Echo{
	e.Pre(middleware.RemoveTrailingSlash())
	return e
}