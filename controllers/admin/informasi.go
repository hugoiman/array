package admin

import (
  "net/http"
  "fmt"
  "array/models/admin"
  "array/structs"
  "github.com/labstack/echo"
  // "encoding/json"
  "time"
)

func ShowAllInformasi(c echo.Context) error {
  status     := c.QueryParam("status")
  session, _ := store.Get(c.Request(), "session")
  slug       := fmt.Sprintf("%v", session.Values["slug"])
  data_admin := admin.GetAdmin(slug)
  informasi  := admin.GetAllInformasi(status)

  data := struct {
    Informasi   structs.Informasi
    Admin       structs.DataAdmin
    Nav         string
    Nav_info    string
  } {
    informasi,
    data_admin,
    "Informasi",
    status,
  }

  // fmt.Printf("%+v\n",data.Informasi.Informasi[0].Waktu)
  return c.Render(http.StatusOK, "informations.html", data)
  // return c.JSON(http.StatusOK,data)
}

func ShowInformasi(c echo.Context) error {
  id_informasi :=  c.Param("id_informasi")
  session, _ := store.Get(c.Request(), "session")
  slug       := fmt.Sprintf("%v", session.Values["slug"])
  data_admin := admin.GetAdmin(slug)
  informasi  := admin.GetInformasi(id_informasi)

  data := struct {
    Informasi   structs.DataInformasi
    Admin       structs.DataAdmin
    Nav         string
  } {
    informasi,
    data_admin,
    "Informasi",
  }

  // fmt.Printf("%+v\n",data)
  return c.Render(http.StatusOK, "information.html", data)
}

func ShowNewInformasi(c echo.Context) error {
  session, _ := store.Get(c.Request(), "session")
  slug       := fmt.Sprintf("%v", session.Values["slug"])
  data_admin := admin.GetAdmin(slug)

  data := struct {
    Admin       structs.DataAdmin
    Nav         string
    Nav_info    string
  } {
    data_admin,
    "Informasi",
    "",
  }

  // fmt.Printf("%+v\n",data)
  return c.Render(http.StatusOK, "create_information.html", data)
}

func CreateInformasi(c echo.Context) error {
  judul       :=  c.FormValue("judul")
  keterangan  :=  c.FormValue("keterangan")
  now         :=  time.Now()
  waktu       :=  now.Format("2006-01-02 15:04:05")
  status      :=  "aktif"

  admin.CreateInformasi(judul, keterangan, waktu, status)

  // decoder   := json.NewDecoder(c.Request().Body)
  // informasi := admin.DataInformasi{}
  //
  // err := decoder.Decode(&informasi)
  // checkErr(err)
  //
  // // admin.CreateInformasi(informasi)
  // fmt.Println(informasi)
  message := "true"
  return c.String(http.StatusOK, message)
}

func DeleteInformasi(c echo.Context) error {
  id_informasi  :=  c.Param("id_informasi")
  admin.DeleteInformasi(id_informasi)
  return nil
}

func CountInformasi(c echo.Context) error {
  semua, aktif, non, draft := admin.CountInformasi()
  data := []string{semua, aktif, non, draft}
  return c.JSON(http.StatusOK, data)
}
