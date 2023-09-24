package router

import (
	"clean-architecture/controller"
	"clean-architecture/logger"
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.Static("/assets", "dist")
	renderer := &TemplateRender{
		templateDir:    "views/",
		layoutTemplate: "layout",
	}
	e.Renderer = renderer
	e.POST("/", uc.Logout)
	e.POST("/login", uc.Login)
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
