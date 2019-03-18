package member

import (
  "fmt"
  "net/http"
  "array/models/member"
  "array/structs"
  "github.com/labstack/echo"
  "crypto/sha1"
  "os"
  "io"
)

func ShowProfil(c echo.Context) error {
  slug := c.Param("slug")
  checkSlug(c, slug)
  data_member    := member.GetMember(slug)

  data := struct {
    Member    structs.DataMember
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

func UpdateFoto(c echo.Context) error {
  session, _ := store.Get(c.Request(), "session")
  id_member  := fmt.Sprintf("%v", session.Values["id_member"])
  slug       := fmt.Sprintf("%v", session.Values["slug"])
  old_foto   := member.GetMember(slug)
  path       := "assets/images/user/"+old_foto.Foto

	file, handler, _ := c.Request().FormFile("file")
  defer file.Close()

  file_name := id_member+"_"+handler.Filename

  dst, _ := os.OpenFile("assets/images/user/"+file_name, os.O_WRONLY|os.O_CREATE, 0666)
  defer dst.Close()

  io.Copy(dst, file)

  os.Remove(path)

  member.UpdateFoto(id_member, file_name)

  return c.Redirect(http.StatusMovedPermanently, "/profil/" + slug)
}
