package controller

import (
	"clean-architecture/logger"
	"clean-architecture/model"
	"clean-architecture/usecase"
	"net/http"

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
	var post model.Post
	if err := c.Bind(&post); err != nil {
		logger.L.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := pc.pu.Create(&post); err != nil {
		logger.L.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.Redirect(http.StatusOK, "/")
}
