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

  e.GET("/informasi/:slug", func(c echo.Context) error{
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
      "views/member/home.html",
      "views/member/head.html", "views/member/header.html","views/member/footer.html",
      )),
    }
    member.ShowInfo(c)
    return nil
  })


  e.GET("/profil/:slug", func(c echo.Context) error{
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
      "views/member/profile.html",
      "views/member/head.html", "views/member/header.html", "views/member/footer.html",
      )),
    }
    member.ShowProfil(c)
    return nil
  })

  e.GET("/administrasi/:slug", func(c echo.Context) error{
    funcs := template.FuncMap{"counter": counter}
    e.Renderer = &Template{ templates: template.Must(template.New("views/member/administrasi.html").Funcs(funcs).ParseFiles(
      "views/member/administrasi.html",
      "views/member/head.html", "views/member/header.html", "views/member/footer.html",
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

  e.GET("/dashboard/:slug", func(c echo.Context) error{
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
      "views/admin/dashboard.html",
      "views/admin/head.html", "views/admin/header.html", "views/admin/footer.html",
      )),
    }
    admin.ShowDashboard(c)
    return nil
  })

  e.GET("/profile/:slug", func(c echo.Context) error{
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
      "views/admin/profile.html",
      "views/admin/head.html", "views/admin/header.html", "views/admin/footer.html",
      )),
    }
    admin.ShowProfil(c)
    return nil
  })

  e.POST("/admin/check-email", admin.CheckEmail)
  e.POST("/admin/update-biodata", admin.UpdateBiodata)
  e.POST("/admin/update-password", admin.UpdatePassword)
  e.POST("/admin/update-foto", admin.UpdatePassword)

  e.GET("/members", func(c echo.Context) error{
    funcs := template.FuncMap{"counter": counter}
    e.Renderer = &Template{ templates: template.Must(template.New("views/member/members.html").Funcs(funcs).ParseFiles(
      "views/admin/members.html",
      "views/admin/head.html", "views/admin/header.html", "views/admin/footer.html",
      )),
    }
    admin.ShowMembers(c)
    return nil
  })

  var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

  echo.NotFoundHandler = func(c echo.Context) error {
    session, _ := store.Get(c.Request(), "session")
    if session.Values["id_member"] != nil {
      e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
        "views/member/error404.html","views/member/head.html","views/member/header.html","views/member/footer.html",
        )),
      }
    } else if session.Values["id_admin"] != nil {
      e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
        "views/admin/err404.html","views/admin/head.html","views/admin/header.html","views/admin/footer.html",
        )),
      }
    } else {
      e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
        "views/guest/index.html","views/guest/head.html",
        )),
      }
    }

    auth.Err404(c)
    return nil
  }

  fmt.Println("server started at :9000")
  e.Logger.Fatal(e.Start(":9000"))
}
