package member

import (
  "net/http"
  "array/models/member"
  "github.com/labstack/echo"
  "fmt"
)

func checkErr(err error){
  if err != nil {
    panic(err.Error())
  }
}

func GetInfo(c echo.Context) error{
  informasi   := member.GetInfo()
  dataMember  := GetMember(c)

  data := struct {
    X member.Informasi
    Y member.DataMember
  } {
    informasi,
    dataMember,
  }

  fmt.Printf("%+v\n",data)
  fmt.Printf("%+v\n",informasi)
  return c.Render(http.StatusOK, "home.html", data)
  // return c.JSON(http.StatusOK, data)
}
