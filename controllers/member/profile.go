package member

import (
  "fmt"
  "net/http"
  "array/models/member"
  "github.com/labstack/echo"
  "crypto/sha1"
)

func ShowProfil(c echo.Context) error {
  slug := c.Param("slug")
  checkSlug(c, slug)
  data_member    := member.GetMember(slug)

  data := struct {
    Member    member.DataMember
    Nav       string
  } {
    data_member,
    "Profil",
  }

  return c.Render(http.StatusOK, "profile.html", data)
    // return c.JSON(http.StatusOK, data)
}

func UpdatePassword(c echo.Context) error{
  id_member   := c.FormValue("id_member")
  newPassword := c.FormValue("password_baru")
  oldPassword := member.GetPassword(id_member)

  var sha = sha1.New()
  sha.Write([]byte(newPassword))
  var encrypted = sha.Sum(nil)
  var encryptedString = fmt.Sprintf("%x", encrypted)

  if encryptedString == oldPassword {
    message := "false"
    return c.String(http.StatusOK, message)
  } else {
    member.UpdatePassword(id_member,encryptedString)
    message := "true"
    return c.String(http.StatusOK, message)
  }
}

// func UpdateFoto(c echo.Context) error {
//   id_member   := c.FormValue("id_member")
//   newPassword := c.FormValue("password_baru")
// }
