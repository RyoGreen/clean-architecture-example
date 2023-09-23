package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(e echo.Context) error {
	return e.Render(http.StatusOK, "top.html", nil)
}

func Login(e echo.Context) error {
	return e.Render(http.StatusOK, "login.html", nil)
}

func Signup(e echo.Context) error {
	return e.Render(http.StatusOK, "signup.html", nil)
}
