package controller

import (
	"clean-architecture/cookie"
	"clean-architecture/logger"
	"clean-architecture/model"
	"clean-architecture/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	Signup(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (u *userController) Signup(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		logger.L.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := u.uu.Signup(user)
	if err != nil {
		logger.L.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

func (u *userController) Login(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		logger.L.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenStr, err := u.uu.Login(user)
	if err != nil {
		logger.L.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.SetCookie(cookie.SetTokenCookie(tokenStr))
	return c.NoContent(http.StatusOK)
}

func (u *userController) Logout(c echo.Context) error {
	c.SetCookie(cookie.DeleteTokenCookie())
	return c.NoContent(http.StatusOK)
}
