package admin

import (
  "fmt"
  "net/http"
  "array/models/admin"
  "array/models/member"
  "github.com/labstack/echo"
  "strconv"
  "strings"
  "encoding/json"
  setRandom "github.com/sethvargo/go-password/password"
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

  slug_member   :=  c.Param("slug")
  data_admin    :=  admin.GetAdmin(slug_admin)
  data_member   :=  member.GetMember(slug_member)

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

func ShowRegistrationMember(c echo.Context) error {
  session, _    :=  store.Get(c.Request(), "session")
  slug          :=  fmt.Sprintf("%v", session.Values["slug"])
  data_admin    :=  admin.GetAdmin(slug)

  data := struct {
    Admin     admin.DataAdmin
    Nav       string
  } {
    data_admin,
    "Member",
  }
  // return c.JSON(http.StatusOK, data)
  return c.Render(http.StatusOK, "registration_member.html", data)
}

func CreateMember(c echo.Context) error {
  decoder   :=  json.NewDecoder(c.Request().Body)
  data      :=  admin.DataMember{}
  err       :=  decoder.Decode(&data)
  checkErr(err)

  data.Password, _ = setRandom.Generate(12, 8, 0, true, true)
  slug      :=  strings.Replace(data.Nama," ","-",-1)
  isUnique  :=  admin.CheckUniqueSlug(slug)

  if isUnique == true {
    data.Slug = slug
    // admin.CreateMember(data)
  } else {
    data.Slug = strings.Join([]string{slug,data.Nik},"-")
    // admin.CreateMember(data)
  }
  fmt.Printf("%+v\n",data)
  return c.String(http.StatusOK,"true")
  // return c.JSON(http.StatusOK, data) // <- blum fix
}

func UpdateMember(c echo.Context) error {
  decoder   :=  json.NewDecoder(c.Request().Body)
  data      :=  admin.DataMember{}
  err       :=  decoder.Decode(&data)
  checkErr(err)

  // admin.UpdateMember(data)
  return c.JSON(http.StatusOK, data) // <- blum fix
}

func DeleteMember(c echo.Context) error {
  id_member, _ := strconv.Atoi(c.Param("id_member"))
  admin.DeleteMember(id_member)
  return c.String(http.StatusOK, "good") // <- blum fix
}
