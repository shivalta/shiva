package baseResponse

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type BaseResponse struct {
	Status   string      `json:"status"`
	Data     interface{} `json:"data"`
	Messages string      `json:"message"`
}

func SuccessResponse(c echo.Context, data interface{}, message string) error {
	response := BaseResponse{}
	response.Status = "success"
	response.Data = data
	response.Messages = message
	return c.JSON(http.StatusOK, response)
}

func ErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Status = "error"
	response.Messages = err.Error()
	return c.JSON(status, response)
}
