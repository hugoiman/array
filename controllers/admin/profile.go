package admin

import (
  "fmt"
  "net/http"
  "array/models/admin"
  "github.com/labstack/echo"
  "crypto/sha1"
  "encoding/json"
)

func ShowProfil(c echo.Context) error {
  slug := c.Param("slug")
  checkSlug(c, slug)
  data_admin  := admin.GetAdmin(slug)

  data := struct {
    Admin    admin.DataAdmin
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

func UpdateBiodata(c echo.Context) error {
  decoder := json.NewDecoder(c.Request().Body)
  biodata := admin.Biodata{}

  err := decoder.Decode(&biodata);
  checkErr(err)
  
  // nama := biodata.Nama
  // fmt.Printf("%+v\n",biodata)
  //
  // biodata.Slug = nama

  fmt.Printf("%+v\n",biodata)

  // admin.UpdateBiodata(biodata.Id_admin, biodata.Nama)

  message := "true"
  return c.String(http.StatusOK, message)
}
