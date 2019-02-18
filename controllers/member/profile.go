package member

import (
  "fmt"
  "net/http"
  "array/models/member"
  "github.com/labstack/echo"
  "crypto/sha1"
)

func ShowProfil(c echo.Context) error {
  slug        := c.Param("slug")
  // dataMember  := GetMember(c, slug)
  dataProfil  := member.GetProfil(slug)

  // fmt.Printf("%+v\n",dataProfil)
  return c.Render(http.StatusOK, "profile.html", dataProfil)
}

func GantiPassword(c echo.Context) error {
  id_member   := c.FormValue("id_member")
  newPassword := c.FormValue("password_baru")
  oldPassword := member.GetPassword(id_member)

  var sha = sha1.New()
  sha.Write([]byte(newPassword))
  var encrypted = sha.Sum(nil)

  var encryptedString = fmt.Sprintf("%x", encrypted)

  if encryptedString == oldPassword {
    fmt.Println("Gagal! Password baru tidak boleh sama dengan password lama.")
  } else {
    fmt.Println("Berhasil mengubah password.")
  }

  return c.JSON(http.StatusOK, encryptedString)
}
