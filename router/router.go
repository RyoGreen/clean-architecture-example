package router

import (
	"clean-architecture/controller"
	"clean-architecture/logger"
	"fmt"
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, pc controller.IPostController, sc controller.ISessionController) *echo.Echo {
	e := echo.New()
	e.Static("/assets", "dist")
	render := &TemplateRender{
		templateDir:    "views/",
		layoutTemplate: "layout",
	}
	e.Renderer = render
	e.HTTPErrorHandler = customErrorHandler
	e.Use(sc.Middleware)
	e.GET("/", pc.ListPosts)
	e.POST("/post", pc.Create)
	e.POST("/logout", uc.Logout)
	e.POST("/login", uc.Login)
	e.GET("/login", uc.IndexLogin)
	e.GET("/signup", uc.IndexSignup)
	e.POST("/signup", uc.Signup)
	return e
}

type TemplateRender struct {
	templateDir    string
	layoutTemplate string
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, e echo.Context) error {
	templates, err := template.ParseFiles(t.templateDir+name, t.templateDir+t.layoutTemplate+".html")
	if err != nil {
		logger.L.Error(err.Error())
		return err
	}
	return templates.ExecuteTemplate(w, t.layoutTemplate, data)
}

func customErrorHandler(err error, c echo.Context) {
	var code int
	var message string

	switch err.(type) {
	case *echo.HTTPError:
		httpErr := err.(*echo.HTTPError)
		code = httpErr.Code
		message = httpErr.Message.(string)
	default:
		code = http.StatusInternalServerError
		message = err.Error()
	}

	var data = map[string]string{
		"Code":    fmt.Sprintf("%v", code),
		"Message": message,
	}
	c.Render(code, "error.html", data)
}

type contextType string

const ContextKey contextType = "current_user"
