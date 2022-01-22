package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shiva/shiva-auth/internal/products"
	"shiva/shiva-auth/utils/baseResponse"
	"shiva/shiva-auth/utils/converter"
)

type Http struct {
	usecase products.Usecase
}

func NewProductsHandler(u products.Usecase) *Http {
	return &Http{
		usecase: u,
	}
}

func (h *Http) GetAll(c echo.Context) error {
	search := c.QueryParam("search")
	key := c.QueryParam("key")
	data, err := h.usecase.GetAll(search, key)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, FromListDomain(data), "berhasil mendapatkan data product products")
}

func (h *Http) GetById(c echo.Context) error {
	id := c.Param("id")
	convId, err := converter.StringToUint(id)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}
	res, err := h.usecase.GetById(convId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, FromDomain(res), "berhasil mendapatkan data")
}

func (h *Http) Create(c echo.Context) error {
	req := new(Request)
	if err := c.Bind(req); err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	res, err := h.usecase.Create(req.ToDomain())
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, FromDomain(res), "data berhasil ditambah!")
}

func (h *Http) Update(c echo.Context) error {
	id := c.Param("id")
	convId, err := converter.StringToUint(id)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}
	tempReq := Request{}
	err = c.Bind(&tempReq)
	if err != nil {
		return err
	}
	req := tempReq.ToDomain()
	req.ID = convId
	res, err := h.usecase.Update(req)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}
	return baseResponse.SuccessResponse(c, FromDomain(res), "data berhasil diupdate!")
}

func (h *Http) Delete(c echo.Context) error {
	id := c.Param("id")
	convId, err := converter.StringToUint(id)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}
	err = h.usecase.Delete(convId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, convId, "data berhasil dihapus!")
}
