package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shiva/shiva-auth/internal/accounts"
	"shiva/shiva-auth/utils/baseResponse"
)

type Http struct {
	usecase accounts.Usecase
}

func NewAccountsHandler(u accounts.Usecase) *Http {
	return &Http{
		usecase: u,
	}
}

func (h *Http) GetAll(c echo.Context) error {
	search := c.QueryParam("search")
	data, err := h.usecase.GetAll(search)
	if err != nil {

	}
	return baseResponse.SuccessResponse(c, FromListDomain(data), "success get all data user")
}

func (h *Http) Create(c echo.Context) error {
	req := Request{}
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	res, err := h.usecase.Create(ToDomain(req))
	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    FromDomain(res),
	})
	//res := h.usecase.Create()
	//res := h.
	//return errors.New("")
}
