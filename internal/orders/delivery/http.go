package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shiva/shiva-auth/cmd/http/middlewares"
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

func (h *Http) GetHistory(c echo.Context) error {
	u := middlewares.CustomContext{Context: c}
	userId, err := u.GetUserId()
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	res, err := h.usecase.GetHistory(userId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, FromDomainToTransactionResponseList(res), "berhasil mendapatkan riwayat transaksi")
}

func (h *Http) CreateVA(c echo.Context) error {
	req := new(RequestPayment)
	if err := c.Bind(req); err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	u := middlewares.CustomContext{Context: c}
	userId, err := u.GetUserId()
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	res, err := h.usecase.CreateVA(req.ProductId, userId, req.BankCode, req.UserValue)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, FromDomainToCreateVAResponse(res), "berhasil melakukan checkout")
}

func (h *Http) PaymentMethod(c echo.Context) error {
	data, err := h.usecase.PaymentChannels()
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, FromDomainPaymentMethodList(data), "berhasil mengambil seluruh payment method")
}

func (h *Http) Checkout(c echo.Context) error {
	req := new(RequestCheckout)
	if err := c.Bind(req); err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	res, err := h.usecase.Checkout(req.UserValue, req.ProductId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, FromDomainToCheckout(res), "berhasil mendapatkan detail checkout")
}
