package member

import (
  "fmt"
  "net/http"
  "array/models/member"
  "github.com/labstack/echo"
)

func ShowAdministrasi(c echo.Context) error {
  session, _  := store.Get(c.Request(), "session")
  id_member   := fmt.Sprintf("%v", session.Values["id_member"])

  dataMember  := member.GetMember(id_member)
  administrasi:= member.GetAdministrasi(id_member)

  data := struct {
    Administrasi   member.Administrasi
    Member         member.DataMember
    Nav            string
  } {
    administrasi,
    dataMember,
    "administrasi",
  }

    // return c.JSON(http.StatusOK, data)
    return c.Render(http.StatusOK, "administrasi.html", data)
}
