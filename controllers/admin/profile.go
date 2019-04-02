package admin

import (
  "fmt"
  "net/http"
  "array/models/admin"
  "array/structs"
  "github.com/labstack/echo"
  "crypto/sha1"
  "encoding/json"
  "os"
  "io"
)

func ShowProfil(c echo.Context) error {
  slug := c.Param("slug")
  checkSlug(c, slug)
  data_admin  := admin.GetAdmin(slug)

  data := struct {
    Admin     structs.DataAdmin
    Nav       string
  } {
    data_admin,
    "Profil",
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

func CheckEmail(c echo.Context) error {
  new_email := c.FormValue("new_email")
  isValid   := admin.CheckEmail(new_email)
  fmt.Println(isValid)
  if isValid == true {
    message := "true"
    return c.String(http.StatusOK, message)
  } else {
    message := "false"
    return c.String(http.StatusOK, message)
  }
}

func UpdateProfil(c echo.Context) error {
  decoder := json.NewDecoder(c.Request().Body)
  data    := structs.DataAdmin{}
  err     := decoder.Decode(&data);
  checkErr(err)

  admin.UpdateProfil(data)

  return c.String(http.StatusOK, "true")
}

func UpdateFoto(c echo.Context) error {
  session, _ := store.Get(c.Request(), "session")
  id_admin   := fmt.Sprintf("%v", session.Values["id_admin"])
  slug       := fmt.Sprintf("%v", session.Values["slug"])
  old_foto   := admin.GetAdmin(slug)
  path       := "assets/images/admin/"+old_foto.Foto

	file, handler, _ := c.Request().FormFile("file")
  defer file.Close()

  file_name := id_admin+"_"+handler.Filename
  dst, _ := os.OpenFile("assets/images/admin/"+file_name, os.O_WRONLY|os.O_CREATE, 0666)
  defer dst.Close()

  io.Copy(dst, file)
  os.Remove(path)

  admin.UpdateFoto(id_admin, file_name)

  return c.Redirect(http.StatusMovedPermanently, "/profile/" + slug)
}
