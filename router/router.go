package router

import (
	"clean-architecture/controller"
	"clean-architecture/logger"
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Static("/assets", "assets")
	renderer := &TemplateRender{
		templateDir:    "views/",
		layoutTemplate: "layout",
	}
	e.Renderer = renderer
	e.GET("/", controller.Index)
	e.GET("/login", controller.Login)
	e.GET("/signup", controller.Signup)
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
