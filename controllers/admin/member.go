package admin

import (
  "fmt"
  "net/http"
  "array/models/admin"
  "array/models/member"
  "github.com/labstack/echo"
)

func ShowMembers(c echo.Context) error {
  session, _    :=  store.Get(c.Request(), "session")
  slug          :=  fmt.Sprintf("%v", session.Values["slug"])
  data_admin    :=  admin.GetAdmin(slug)
  data_members  :=  admin.GetMembers()

  data := struct {
    Admin     admin.DataAdmin
    Member    admin.Member
    Nav       string
  } {
    data_admin,
    data_members,
    "Member",
  }
  // return c.JSON(http.StatusOK, data)
  return c.Render(http.StatusOK, "members.html", data)
}

func ShowProfilMember(c echo.Context) error {
  session, _    :=  store.Get(c.Request(), "session")
  slug_admin    :=  fmt.Sprintf("%v", session.Values["slug"])

  slug_member   := c.Param("slug")
  data_admin    := admin.GetAdmin(slug_admin)
  data_member   := member.GetMember(slug_member)

  data := struct {
    Admin     admin.DataAdmin
    Member    member.DataMember
    Nav       string
  } {
    data_admin,
    data_member,
    "Member",
  }

  return c.Render(http.StatusOK, "profile_member.html", data)
}
