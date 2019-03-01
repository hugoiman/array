package admin

import (
  "fmt"
  "net/http"
  "array/models/admin"
  "github.com/labstack/echo"
  "crypto/sha1"
)

func ShowProfil(c echo.Context) error {
  session, _  := store.Get(c.Request(), "session")
  id_admin    := fmt.Sprintf("%v", session.Values["id_admin"])
  data_admin  := admin.GetAdmin(id_admin)

  data := struct {
    Admin    admin.DataAdmin
    Nav       string
  } {
    data_admin,
    "profil",
  }

  return c.Render(http.StatusOK, "profile.html", data)
}

func UpdatePassword(c echo.Context) error{
  id_admin    := c.FormValue("id_admin")
  newPassword := c.FormValue("password_baru")
  oldPassword := admin.GetPassword(id_admin)

  var sha = sha1.New()
  sha.Write([]byte(newPassword))
  var encrypted = sha.Sum(nil)
  var encryptedString = fmt.Sprintf("%x", encrypted)

  if encryptedString == oldPassword {
    message := "false"
    return c.String(http.StatusOK, message)
  } else {
    admin.UpdatePassword(id_admin,encryptedString)
    message := "true"
    return c.String(http.StatusOK, message)
  }
}
