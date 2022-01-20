package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shiva/shiva-auth/internal/accounts"
	"shiva/shiva-auth/utils/baseResponse"
	"shiva/shiva-auth/utils/converter"
)

type Http struct {
	usecase accounts.Usecase
}

func NewAccountsHandler(u accounts.Usecase) *Http {
	return &Http{
		usecase: u,
	}
}

func (h *Http) Login(c echo.Context) error {
	req := RequestLogin{}
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	res, err := h.usecase.Login(req.Email, req.Password)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, res, "login successfuly!")
}

func (h *Http) GetAll(c echo.Context) error {
	search := c.QueryParam("search")
	data, err := h.usecase.GetAll(search)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, FromListDomain(data), "success get all data user")
}

func (h *Http) Create(c echo.Context) error {
	req := Request{}
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	res, err := h.usecase.Create(req.ToDomain())
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, FromDomain(res), "user has been registered!")
}

func (h *Http) Update(c echo.Context) error {
	userId := c.Param("userId")
	convUserId, err := converter.StringToUint(userId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}
	tempReq := Request{}
	err = c.Bind(&tempReq)
	if err != nil {
		return err
	}
	req := tempReq.ToDomain()
	req.ID = convUserId
	res, err := h.usecase.Update(req)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}
	return baseResponse.SuccessResponse(c, FromDomain(res), "update successfuly!")
}

func (h *Http) Delete(c echo.Context) error {
	userId := c.Param("userId")
	convUserId, err := converter.StringToUint(userId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}
	err = h.usecase.Delete(convUserId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, convUserId, "delete successfuly")
}

func (h *Http) GetById(c echo.Context) error {
	userId := c.Param("userId")
	convUserId, err := converter.StringToUint(userId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}
	res, err := h.usecase.GetById(convUserId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, FromDomain(res), "get data successfuly")
}

func (h *Http) Verify(c echo.Context) error {
	req := RequestVerify{}
	err := c.Bind(&req)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}
	v, err := h.usecase.Verify(req.Email, req.Verify)
	if err != nil {
		return err
	}
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return baseResponse.SuccessResponse(c, FromDomain(v), "user has been verified!")
}
