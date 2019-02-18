package member

import (
  "array/models/member"
  "github.com/labstack/echo"
)

func GetMember(c echo.Context) member.DataMember {
  id_member := "13"
  dataMember := member.GetMember(id_member)
  return dataMember
}
