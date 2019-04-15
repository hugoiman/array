package admin

import (
  "fmt"
  "net/http"
  "array/models/admin"
  "array/structs"
  "github.com/labstack/echo"
  "strings"
  // "time"
)

func ShowAdministrasi(c echo.Context) error {
  periode       :=  c.Param("periode")
  session, _    :=  store.Get(c.Request(), "session")
  slug          :=  fmt.Sprintf("%v", session.Values["slug"])
  data_admin    :=  admin.GetAdmin(slug)
  data_administrasi  :=  admin.GetAdministrasi()

  str_periode   :=  strings.Split(periode, "-")
  fmt.Println(str_periode)
  month         :=  str_periode[0]
  year          :=  str_periode[1]
  fmt.Println(month)
  fmt.Println(year)

  data := struct {
    Admin           structs.DataAdmin
    Member          structs.Member
    Nav             string
  } {
    data_admin,
    data_administrasi,
    "Administrasi",
  }
  // return c.JSON(http.StatusOK, data.Member.Member[0].Administrasi)
  // return c.JSON(http.StatusOK, data)
  return c.Render(http.StatusOK, "administration.html", data)
}

// func CreatePembayaran(c echo.Context) error {
//   id_member         := c.FormValue("id_member")
//   check_in          := c.FormValue("check_in")
//   check_out         := c.FormValue("check_out")
//   jumlah_pembayaran := c.FormValue("jumlah_pembayaran")
//   total_pembayaran  := c.FormValue("total_pembayaran")
//
//   d, m, y
//
//   days := t2.Sub(t1).Hours() / 24
//   fmt.Println(days) // 366
// }
//
// getDay(c echo.Context) error{}
