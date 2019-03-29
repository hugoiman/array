package admin

import (
  "net/http"
  "fmt"
  "array/models/admin"
  "array/structs"
  "github.com/labstack/echo"
)

func ShowAllLokasi(c echo.Context) error {
  session, _  := store.Get(c.Request(), "session")
  slug        := fmt.Sprintf("%v", session.Values["slug"])
  data_admin  := admin.GetAdmin(slug)
  data_lokasi := admin.GetAllLokasi()

  data := struct {
    Lokasi      structs.Lokasi
    Admin       structs.DataAdmin
    Nav         string
    Nav2        string
  } {
    data_lokasi,
    data_admin,
    "Pengaturan",
    "Lokasi",
  }

  return c.Render(http.StatusOK, "locations.html", data)
}

func CreateLokasi(c echo.Context) error {
  cabang     := c.FormValue("cabang")
  alamat     := c.FormValue("alamat")

  admin.CreateLokasi(cabang,alamat)

  return nil
}

func UpdateLokasi(c echo.Context) error  {
  id_lokasi  := c.FormValue("id_lokasi")
  cabang     := c.FormValue("cabang")
  alamat     := c.FormValue("alamat")

  admin.UpdateLokasi(id_lokasi, cabang, alamat)

  return nil
}

func DeleteLokasi(c echo.Context) error {
  id_lokasi := c.Param("id_lokasi")

  admin.DeleteLokasi(id_lokasi)
  fmt.Println(id_lokasi)

  return nil
}
