package member

import (
  "fmt"
  "net/http"
  "array/models/member"
  "array/structs"
  "github.com/labstack/echo"
)

func ShowAdministrasi(c echo.Context) error {
  slug := c.Param("slug")
  checkSlug(c, slug)
  session, _      :=  store.Get(c.Request(), "session")
  id_member       :=  fmt.Sprintf("%v", session.Values["id_member"])
  data_member     :=  member.GetMember(slug)
  administrasi    :=  member.GetAdministrasi(id_member)
  infoTagihan     :=  member.GetDetailTagihan(id_member)
  status, tagihan :=  member.GetTagihan(id_member)

  data := struct {
    Administrasi   structs.Administrasi
    Member         structs.DataMember
    InfoTagihan    structs.InfoTagihan
    Status         string
    Tagihan        int
    Nav            string
  } {
    administrasi,
    data_member,
    infoTagihan,
    status,
    tagihan,
    "Administrasi",
  }

  return c.Render(http.StatusOK, "administrasi.html", data)
  // x := administrasi.Administrasi[0].Check_in
  // var y = x.Format("Mon, 02 Jan 2006")
  // fmt.Println(x)

  // return c.JSON(http.StatusOK, data)
}
