package main

import (
  "net/http"
  "fmt"
  "html/template"
  "io"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/gorilla/context"

  "array/controllers/auth"
  "array/controllers/member"
  "array/controllers/admin"

  "github.com/gorilla/sessions"
  "os"
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

func counter(x, y int) int {
  return x + y
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

  //  Auth
  e.GET("/", func(c echo.Context) error{
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
      "views/guest/index.html",
      "views/guest/head.html",
      )),
    }
    auth.Index(c)
    return nil
  })

  e.POST("/login", auth.Login)

  //  Member
  e.GET("/logout", func(c echo.Context) error{
    var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles("views/guest/index.html")), }
    session, _ := store.Get(c.Request(), "session")
    session.Options.MaxAge = -1
    session.Save(c.Request(), c.Response())

    return c.Redirect(http.StatusMovedPermanently, "/")
  })

  e.GET("/informasi", func(c echo.Context) error{
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
      "views/member/home.html",
      "views/member/head.html",
      "views/member/header.html",
      "views/member/footer.html",
      )),
    }
    e.GET("/profil", func(c echo.Context) error{
    member.ShowInfo(c)
    return nil
  })

    e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
      "views/member/profile.html",
      "views/member/head.html",
      "views/member/header.html",
      "views/member/footer.html",
      )),
    }
    member.ShowProfil(c)
    return nil
  })

  e.GET("/administrasi", func(c echo.Context) error{
    funcs := template.FuncMap{"counter": counter}
    e.Renderer = &Template{ templates: template.Must(template.New("views/member/administrasi.html").Funcs(funcs).ParseFiles(
      "views/member/administrasi.html",
      "views/member/head.html",
      "views/member/header.html",
      "views/member/footer.html",
      )),
    }
    member.ShowAdministrasi(c)
    return nil
  })

  e.POST("/update-password", member.UpdatePassword)

  e.POST("/update-foto", member.UpdatePassword)

  //  Admin
  e.GET("/admin", func(c echo.Context) error {
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
      "views/guest/indexadmin.html",
      "views/guest/head.html",
      )),
    }
    auth.IndexAdmin(c)
    return nil
  })

  e.POST("/login-admin", auth.LoginAdmin)

  e.GET("/dashboard", func(c echo.Context) error{
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
      "views/admin/dashboard.html",
      "views/admin/head.html",
      "views/admin/header.html",
      "views/admin/footer.html",
      )),
    }
    admin.ShowDashboard(c)
    return nil
  })

  echo.NotFoundHandler = func(c echo.Context) error {
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
      "views/admin/err404.html","views/admin/head.html","views/admin/header.html","views/admin/footer.html",
      "views/member/error404.html","views/member/head.html","views/member/header.html","views/member/footer.html",
      )),
    }
    auth.Err404(c)
    // user_input  :=  c.Request().URL    //  http.URL
    // msg, _      :=   "s", user_input
    // render your 404 page
    // return c.String(http.StatusNotFound, msg)
    return nil
  }

  fmt.Println("server started at :9000")
  e.Logger.Fatal(e.Start(":9000"))
}
