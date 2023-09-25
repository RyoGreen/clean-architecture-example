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
		return err
	}
	if _, err := u.uu.Signup(user); err != nil {
		logger.L.Error(err.Error())
		return err
	}
	return c.Redirect(http.StatusFound, "/login")
}

func (u *userController) Login(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		logger.L.Error(err.Error())
		return err
	}
	tokenStr, err := u.uu.Login(user)
	if err != nil {
		logger.L.Error(err.Error())
		return err
	}
	c.SetCookie(cookie.SetTokenCookie(tokenStr))
	return c.Redirect(http.StatusFound, "/")
}

func (u *userController) Logout(c echo.Context) error {
	c.SetCookie(cookie.DeleteTokenCookie())
	return c.Redirect(http.StatusFound, "/")
}
