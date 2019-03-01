package auth

import (
  "net/http"
  "fmt"
	"array/models/auth"
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
    return c.Redirect(http.StatusMovedPermanently, "/dashboard")
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
    dataSession := auth.SessionAdmin{}

    dataSession.Id_admin = result.Id_admin
    dataSession.Nama = result.Nama
    dataSession.Slug = result.Slug

    message := "true"
    return c.String(http.StatusOK, message)
	} else {
    message := "false"
    return c.String(http.StatusOK, message)
	}
}

func SetSessionAdmin(c echo.Context, data auth.SessionAdmin) {
  session, _ := store.Get(c.Request(), "session")
  session.Values["id_admin"] = data.Id_admin
  session.Values["nama"] = data.Nama
  session.Values["slug"] = data.Slug
  session.Save(c.Request(), c.Response())
}
