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
    Administrasi    structs.Administrasi
    Nav             string
  } {
    data_admin,
    data_administrasi,
    "Administrasi",
  }
  // return c.JSON(http.StatusOK, data)
  return c.Render(http.StatusOK, "administration.html", data)
}
