package controller

import (
	"clean-architecture/cookie"
	"clean-architecture/logger"
	"clean-architecture/usecase"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ISessionController interface {
	Middleware(next echo.HandlerFunc) echo.HandlerFunc
}

type sessionController struct {
	uu usecase.ISessionUsecase
}

func NewSessionController(su usecase.ISessionUsecase) ISessionController {
	return &sessionController{su}
}

type contextType string

const contextKey contextType = "current_user"

func (sc *sessionController) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		getCookie, err := c.Cookie(cookie.SID)
		if err != nil {
			if err == http.ErrNoCookie {
				if err := next(c); err != nil {
					c.Error(err)
				}
				return nil
			}
			logger.L.Error(err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		userSid, err := sc.uu.GetUserSid(getCookie)
		if err != nil {
			logger.L.Info(err.Error())
			return nil
		}
		user, err := sc.uu.GetSessoionByID(userSid)
		if err != nil {
			logger.L.Info(err.Error())
			return nil
		}
		ctx := context.WithValue(c.Request().Context(), contextKey, user)
		c.SetRequest(c.Request().WithContext(ctx))
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
