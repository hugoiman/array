package member

import (
  "net/http"
  "fmt"
  "array/models/member"
  "github.com/labstack/echo"
)

func checkErr(err error){
  if err != nil {
    panic(err.Error())
  }
}

func GetInfo(c echo.Context) error{
  informasi := member.GetInfo()
  // member    := GetMember(c)
  fmt.Printf("%+v\n",informasi)
  return c.Render(http.StatusOK, "home.html", informasi)
}
