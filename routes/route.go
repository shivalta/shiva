package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RouteV1(e *echo.Echo) {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/")
}

func handlerIndex(writer http.ResponseWriter, request *http.Request) {

}
