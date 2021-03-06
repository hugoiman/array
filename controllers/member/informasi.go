package member

import (
  "net/http"
  "array/models/member"
  "array/structs"
  "github.com/labstack/echo"
  "github.com/gorilla/sessions"
  "os"
  // "reflect"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func ShowAllInformasi(c echo.Context) error{
  slug := c.Param("slug")
  checkSlug(c, slug)
  data_member := member.GetMember(slug)
  informasi   := member.GetAllInformasi()
  // fmt.Println(reflect.TypeOf(id_member))
  data := struct {
    Informasi structs.Informasi
    Member    structs.DataMember
    Nav       string
  } {
    informasi,
    data_member,
    "Informasi",
  }

  // fmt.Printf("%+v\n",data)
  return c.Render(http.StatusOK, "home.html", data)
}
