package admin

import (
  "fmt"
  "net/http"
  "array/models/admin"
  "array/structs"
  "github.com/labstack/echo"
  "strings"
  "strconv"
  "encoding/json"
  // "time"
)

var month = [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

func ShowAdministrasi(c echo.Context) error {
  from          :=  c.QueryParam("from")
  to            :=  c.QueryParam("to")
  session, _    :=  store.Get(c.Request(), "session")
  slug          :=  fmt.Sprintf("%v", session.Values["slug"])
  data_admin    :=  admin.GetAdmin(slug)

  data_form    :=  []string{from, to}
  split_from   :=  strings.Split(from, " ")
  split_to     :=  strings.Split(to, " ")

  var startDate, endDate string
  for index, element := range month{
    if split_from[0] == element {
      startMonth := strconv.Itoa(index+1)
      startDate = strings.Join([]string{split_from[1],startMonth,"01"},"-")
    }
    if split_to[0] == element {
      endMonth := strconv.Itoa(index+1)
      endDate = strings.Join([]string{split_to[1],endMonth,"31"},"-")
    }
  }
  fmt.Println(startDate, endDate)

  data_administrasi  :=  admin.GetAdministrasi(startDate, endDate)

  data := struct {
    Admin           structs.DataAdmin
    Member          structs.Member
    FormValue       []string
    Nav             string
  } {
    data_admin,
    data_administrasi,
    data_form,
    "Administrasi",
  }
  // return c.JSON(http.StatusOK, data.Member.Member[0].Administrasi)
  // return c.JSON(http.StatusOK, data)
  return c.Render(http.StatusOK, "administration.html", data)
}

func DeleteAdministrasi(c echo.Context) error  {
  decoder := json.NewDecoder(c.Request().Body)
  data    := struct {
    Id_administrasi   []string  `json:"id_administrasi"`
  }{}

  err     := decoder.Decode(&data);
  checkErr(err)

  for i := 0; i < len(data.Id_administrasi); i++ {
    admin.DeleteAdministrasi(data.Id_administrasi[i])
  }
  return nil
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
