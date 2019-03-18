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

  //  error404
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
  e.POST("/reset_password", auth.ResetPassword)

  e.GET("/reset", func(c echo.Context) error {
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
      "views/guest/reset_password.html",
      "views/guest/head.html",
      )),
    }
    data := struct {
      o string
    }{
      "1",
    }
    return c.Render(http.StatusOK, "reset_password.html", data)
  })

  //  Member
  e.GET("/logout", func(c echo.Context) error{
    var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles("views/guest/index.html")), }
    session, _ := store.Get(c.Request(), "session")
    session.Options = &sessions.Options {
      MaxAge: -1,
    }
    session.Save(c.Request(), c.Response())

    return c.Redirect(http.StatusMovedPermanently, "/")
  })

  e.GET("/informasi/:slug", func(c echo.Context) error{
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
      "views/member/home.html",
      "views/member/head.html", "views/member/header.html","views/member/footer.html",
      )),
    }
    member.ShowAllInformasi(c)
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
  e.POST("/update-foto", member.UpdatePassword)       // blom

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
  e.POST("/admin/update-foto", admin.UpdateFoto)            // blom
  //
  e.POST("/admin/create-member", admin.CreateMember)        // blom
  // e.POST("/admin/update-member", admin.UpdateMember)        // blom
  e.DELETE("/admin/delete-member/:id_member", admin.DeleteMember)  // blom

  e.POST("/admin/create-informasi", admin.CreateInformasi)
  e.GET("/admin/delete-informasi/:id_informasi", admin.DeleteInformasi)

  e.GET("/members", func(c echo.Context) error {
    funcs := template.FuncMap{"counter": counter}
    e.Renderer = &Template{ templates: template.Must(template.New("views/admin/members.html").Funcs(funcs).ParseFiles(
      "views/admin/members.html",
      "views/admin/head.html", "views/admin/header.html", "views/admin/footer.html",
      )),
    }
    admin.ShowMembers(c)
    return nil
  })

  e.GET("/member/:slug", func(c echo.Context) error {
    funcs := template.FuncMap{"counter": counter}
    e.Renderer = &Template{ templates: template.Must(template.New("views/admin/profile_member.html").Funcs(funcs).ParseFiles(
      "views/admin/profile_member.html",
      "views/admin/head.html", "views/admin/header.html", "views/admin/footer.html",
      )),
    }
    admin.ShowProfilMember(c)
    return nil
  })

  e.GET("/registrasi", func(c echo.Context) error{
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
      "views/admin/registration_member.html",
      "views/admin/head.html", "views/admin/header.html", "views/admin/footer.html",
      )),
    }
    admin.ShowRegistrationMember(c)
    return nil
  })

  e.GET("/administration/:periode", func(c echo.Context) error{
    funcs := template.FuncMap{"counter": counter}
    e.Renderer = &Template{ templates: template.Must(template.New("views/admin/administration.html").Funcs(funcs).ParseFiles(
      "views/admin/administration.html",
      "views/admin/head.html", "views/admin/header.html", "views/admin/footer.html",
      )),
    }
    admin.ShowAdministrasi(c)
    return nil
  })

  e.GET("/information/status/:status", func(c echo.Context) error{
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
      "views/admin/informations.html",
      "views/admin/head.html", "views/admin/header.html", "views/admin/footer.html",
      )),
    }
    admin.ShowAllInformasi(c)
    return nil
  })

  e.GET("/information/:id_informasi", func(c echo.Context) error{
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
      "views/admin/information.html",
      "views/admin/head.html", "views/admin/header.html", "views/admin/footer.html",
      )),
    }
    admin.ShowInformasi(c)
    return nil
  })

  e.GET("/new-information", func(c echo.Context) error{
    e.Renderer = &Template{ templates: template.Must(template.ParseFiles(
      "views/admin/create_information.html",
      "views/admin/head.html", "views/admin/header.html", "views/admin/footer.html",
      )),
    }
    admin.ShowNewInformasi(c)
    return nil
  })

  e.POST("/count-information", admin.CountInformasi)

  fmt.Println("server started at :9000")
  e.Logger.Fatal(e.Start(":9000"))
}
