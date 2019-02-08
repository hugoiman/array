package main

import (
  "fmt"
  "html/template"
  "io"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/gorilla/context"

  "array/controllers/auth"
  // "array/controllers/member"
)

type Template struct {
  templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
  // Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
  return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
  e := echo.New()

  e.Static("/static","assets")

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())
  e.Use(echo.WrapMiddleware(context.ClearHandler))

  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
  }))

  e.Renderer = &Template{
    templates: template.Must(template.ParseGlob("views/guest/*.html")),
  }

  //  auth
  e.GET("/", auth.Index)
  e.POST("/login", auth.Login)

  //  member
  e.GET("/logout", auth.Logout)

  fmt.Println("server started at :9000")
  e.Logger.Fatal(e.Start(":9000"))
}
