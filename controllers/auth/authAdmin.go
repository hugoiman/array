package auth

import (
  "net/http"
  "fmt"
	"array/models/auth"
  "array/structs"
  "github.com/labstack/echo"
  "crypto/sha1"
  // "github.com/gorilla/sessions"
  // "os"
)

func IndexAdmin(c echo.Context) error{
  session := CheckSession(c)
  data := struct {
    o string
  }{
    "1",
  }

  if session == false {
    return c.Render(http.StatusOK, "indexadmin.html", data)
  } else {
    session, _  :=  store.Get(c.Request(), "session")
    slug        :=  fmt.Sprintf("%v", session.Values["slug"])
    if session.Values["id_member"] != nil {
      return c.Redirect(http.StatusMovedPermanently, "/informasi/" + slug)
    } else if session.Values["id_admin"] != nil {
      return c.Redirect(http.StatusMovedPermanently, "/information/status/aktif")
    } else {
      return c.Redirect(http.StatusMovedPermanently, "/dashboard")
    }
  }
}

func LoginAdmin(c echo.Context) error{
	email        := c.FormValue("email")
  password     := c.FormValue("password")

  var sha = sha1.New()
  sha.Write([]byte(password))
  var encrypted = sha.Sum(nil)
  var encryptedString = fmt.Sprintf("%x", encrypted)
  authentic       := auth.AuthAdmin(email, encryptedString)

  if authentic == true {
    result       := auth.GetSessionAdmin(email)
    SetSessionAdmin(c, result)
    dataSession := structs.SessionAdmin{}

    dataSession.Id_admin = result.Id_admin
    dataSession.Nama = result.Nama
    dataSession.Slug = result.Slug

    return c.String(http.StatusOK, result.Slug)
	} else {
    message := "false"
    return c.String(http.StatusOK, message)
	}
}

func SetSessionAdmin(c echo.Context, data structs.SessionAdmin) {
  session, _ := store.Get(c.Request(), "session")
  session.Values["id_admin"] = data.Id_admin
  session.Values["nama"] = data.Nama
  session.Values["slug"] = data.Slug
  session.Save(c.Request(), c.Response())
}
