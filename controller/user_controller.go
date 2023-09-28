package controller

import (
	"clean-architecture/cookie"
	"clean-architecture/logger"
	"clean-architecture/model"
	"clean-architecture/usecase"
	"crypto/rand"
	"encoding/base64"
	"io"
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
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.Redirect(http.StatusFound, "/login")
}

func (u *userController) Login(c echo.Context) error {
	if err := c.Request().ParseForm(); err != nil {
		logger.L.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	userID, err := u.uu.Login(c.FormValue("email"), c.FormValue("password"))
	if err != nil {
		logger.L.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	val, err := issuesSIDVal()
	if err != nil {
		logger.L.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	now := time.Now()
	session := &model.Session{
		SID:       val,
		UserID:    userID,
		Expired:   now.Add(time.Hour * 24),
		CreatedAt: now,
	}
	if err := u.uu.CreateSession(session); err != nil {
		logger.L.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	setCookie, err := u.uu.SetSID(val)
	if err != nil {
		logger.L.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.SetCookie(setCookie)
	return c.Redirect(http.StatusFound, "/")
}

func (u *userController) Logout(c echo.Context) error {
	c.SetCookie(cookie.DelSID())
	return c.Redirect(http.StatusFound, "/")
}

func (u *userController) IndexSignup(c echo.Context) error {
	return c.Render(http.StatusOK, "signup.html", nil)
}

func (u *userController) IndexLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

func issuesSIDVal() (string, error) {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
