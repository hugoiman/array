package member

import (
  "fmt"
  "net/http"
  "array/models/member"
  "github.com/labstack/echo"
  // "time"
)

func ShowAdministrasi(c echo.Context) error {
  session, _      :=  store.Get(c.Request(), "session")
  id_member       :=  fmt.Sprintf("%v", session.Values["id_member"])

  dataMember      :=  member.GetMember(id_member)
  administrasi    :=  member.GetAdministrasi(id_member)
  infoTagihan     :=  member.GetDetailTagihan(id_member)
  status, tagihan :=  member.GetTagihan(id_member)

  data := struct {
    Administrasi   member.Administrasi
    Member         member.DataMember
    InfoTagihan    member.InfoTagihan
    Status         string
    Tagihan        int
    Nav            string
  } {
    administrasi,
    dataMember,
    infoTagihan,
    status,
    tagihan,
    "administrasi",
  }
  // x := administrasi.Administrasi[0].Check_in
  // var y = x.Format("Mon, 02 Jan 2006")
  // fmt.Println(x)

  // return c.JSON(http.StatusOK, data)
  return c.Render(http.StatusOK, "administrasi.html", data)
}
