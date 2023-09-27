package controller

import (
	"clean-architecture/cookie"
	"clean-architecture/logger"
	"clean-architecture/model"
	"clean-architecture/usecase"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	Signup(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	IndexSignup(c echo.Context) error
	IndexLogin(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (u *userController) Signup(c echo.Context) error {
	if err := c.Request().ParseForm(); err != nil {
		logger.L.Error(err.Error())
		return err
	}
	now := time.Now()
	var user = model.User{
		Password:  c.Request().FormValue("password"),
		Email:     c.Request().FormValue("email"),
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := u.uu.Signup(user); err != nil {
		logger.L.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Bad Request")
	}
	return c.Redirect(http.StatusFound, "/login")
}

func (u *userController) Login(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		logger.L.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	tokenStr, err := u.uu.Login(user)
	if err != nil {
		logger.L.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Bad Request")
	}
	c.SetCookie(cookie.SetTokenCookie(tokenStr))
	return c.Redirect(http.StatusFound, "/")
}

func (u *userController) Logout(c echo.Context) error {
	c.SetCookie(cookie.DeleteTokenCookie())
	return c.Redirect(http.StatusFound, "/")
}

func (u *userController) IndexSignup(c echo.Context) error {
	return c.Render(http.StatusOK, "signup.html", nil)
}

func (u *userController) IndexLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}
