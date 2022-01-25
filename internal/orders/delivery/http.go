package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shiva/shiva-auth/internal/orders"
	"shiva/shiva-auth/utils/baseResponse"
)

type Http struct {
	usecase orders.Usecase
}

func NewOrdersHandler(u orders.Usecase) *Http {
	return &Http{
		usecase: u,
	}
}

func (h *Http) CheckoutPulsa(c echo.Context) error {
	req := new(RequestCheckout)
	if err := c.Bind(req); err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	res, err := h.usecase.CheckoutPulsa(req.UserValue, req.ProductId, false)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, FromDomainToCheckout(res), "berhasil melakukan checkout")
}

func (h *Http) CheckoutListrik(c echo.Context) error {
	req := new(RequestCheckout)
	if err := c.Bind(req); err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	res, err := h.usecase.CheckoutListrik(req.UserValue, req.ProductId, false)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, FromDomainToCheckout(res), "berhasil melakukan checkout")
}

func (h *Http) CheckoutPDAM(c echo.Context) error {
	req := new(RequestCheckout)
	if err := c.Bind(req); err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	res, err := h.usecase.CheckoutPDAM(req.UserValue, req.ProductId, false)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, FromDomainToCheckout(res), "berhasil melakukan checkout")
}

func (h *Http) CheckoutPulsaWithLogin(c echo.Context) error {
	req := new(RequestCheckout)
	if err := c.Bind(req); err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	res, err := h.usecase.CheckoutPulsa(req.UserValue, req.ProductId, true)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, FromDomainToCheckout(res), "berhasil melakukan checkout")
}

func (h *Http) CheckoutListrikWithLogin(c echo.Context) error {
	req := new(RequestCheckout)
	if err := c.Bind(req); err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	res, err := h.usecase.CheckoutListrik(req.UserValue, req.ProductId, true)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, FromDomainToCheckout(res), "berhasil melakukan checkout")
}

func (h *Http) CheckoutPDAMWithLogin(c echo.Context) error {
	req := new(RequestCheckout)
	if err := c.Bind(req); err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	res, err := h.usecase.CheckoutPDAM(req.UserValue, req.ProductId, true)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, FromDomainToCheckout(res), "berhasil melakukan checkout")
}
