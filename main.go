package main

import (
  // "net/http"
  "fmt"
  "html/template"
  "io"
  //
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/gorilla/context"

  "array/controllers/auth"
  "array/controllers/member"
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

  // e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
  //   Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
  // }))
  e.Use(middleware.Recover())
  e.Use(echo.WrapMiddleware(context.ClearHandler))

  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
  }))

  //  auth
  e.GET("/", func(c echo.Context) error{
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles("views/guest/index.html")), }
    auth.Index(c)
    return nil
  })
  
  e.POST("/login", func(c echo.Context) error{
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
      "views/guest/index.html",
      "views/member/home.html",
      "views/member/header.html",
      "views/member/footer.html",
      )),
    }
    auth.Login(c)
    return nil
  })

  //  member
  e.GET("/logout", func(c echo.Context) error{
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles("views/guest/index.html")), }
    auth.Logout(c)
    return nil
  })

  e.GET("/informasi", func(c echo.Context) error{
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
      "views/member/home.html",
      "views/member/header.html",
      "views/member/footer.html",
      )),
    }
    member.GetInfo(c)
    return nil
  })

  fmt.Println("server started at :9000")
  e.Logger.Fatal(e.Start(":9000"))
}
