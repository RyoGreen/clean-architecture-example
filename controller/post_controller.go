package controller

import (
	"clean-architecture/logger"
	"clean-architecture/model"
	"clean-architecture/usecase"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type IPostController interface {
	ListPosts(c echo.Context) error
	Create(c echo.Context) error
}

type postController struct {
	pu usecase.IPostUsecase
}

func NewPostController(pu usecase.IPostUsecase) IPostController {
	return &postController{pu}
}

func (pc postController) ListPosts(c echo.Context) error {
	posts, err := pc.pu.ListsPosts()
	if err != nil {
		logger.L.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, "index.html", posts)
}

func (pc postController) Create(c echo.Context) error {
	if err := c.Request().ParseForm(); err != nil {
		logger.L.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	now := time.Now()
	var post = model.Post{
		Content:   c.FormValue("content"),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    1,
	}
	if err := pc.pu.Create(&post); err != nil {
		logger.L.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.Redirect(http.StatusFound, "/")
}
