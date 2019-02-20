package member

import (
  "net/http"
  "array/models/member"
  "github.com/labstack/echo"
  "fmt"
  "github.com/gorilla/sessions"
  "os"
  // "reflect"
)

func checkErr(err error){
  if err != nil {
    panic(err.Error())
  }
}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func ShowInfo(c echo.Context) error{
  session, _  := store.Get(c.Request(), "session")
  id_member   := fmt.Sprintf("%v", session.Values["id_member"])
  // fmt.Println(reflect.TypeOf(id_member))
  informasi   := member.GetInfo()
  dataMember  := member.GetMember(id_member)

  data := struct {
    Informasi member.Informasi
    Member    member.DataMember
    Nav       string
  } {
    informasi,
    dataMember,
    "informasi",
  }

  // fmt.Printf("%+v\n",data)
  return c.Render(http.StatusOK, "home.html", data)
}
