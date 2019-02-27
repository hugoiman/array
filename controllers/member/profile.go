package member

import (
  "fmt"
  "net/http"
  "array/models/member"
  "github.com/labstack/echo"
  "crypto/sha1"
)

func ShowProfil(c echo.Context) error {
  session, _  := store.Get(c.Request(), "session")
  id_member   := fmt.Sprintf("%v", session.Values["id_member"])

  dataMember  := member.GetMember(id_member)
  // fmt.Printf("%+v\n",dataMember)

  data := struct {
    Member    member.DataMember
    Nav       string
  } {
    dataMember,
    "profil",
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
