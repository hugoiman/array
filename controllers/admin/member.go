package admin

import (
  "fmt"
  "net/http"
  "array/models/admin"
  "array/models/member"
  "array/structs"
  "github.com/labstack/echo"
  "strconv"
  // "strings"
  "encoding/json"
  // "crypto/sha1"
  // setRandom "github.com/sethvargo/go-password/password"
  "os"
  "io"
)

func ShowMembers(c echo.Context) error {
  session, _    :=  store.Get(c.Request(), "session")
  slug          :=  fmt.Sprintf("%v", session.Values["slug"])
  data_admin    :=  admin.GetAdmin(slug)
  data_members  :=  admin.GetMembers()
  total, aktif, non_aktif, selesai :=  admin.GetCountStatusMember()
  status := []string{total, aktif, non_aktif, selesai}

  data := struct {
    Admin     structs.DataAdmin
    Member    structs.Member
    Status    []string
    Nav       string
  } {
    data_admin,
    data_members,
    status,
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
    Admin     structs.DataAdmin
    Member    structs.DataMember
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
  data_lokasi   :=  admin.GetAllLokasi()

  data := struct {
    Admin     structs.DataAdmin
    Lokasi    structs.Lokasi
    Nav       string
  } {
    data_admin,
    data_lokasi,
    "Member",
  }
  return c.Render(http.StatusOK, "registration_member.html", data)
}

func CreateMember(c echo.Context) error {
  fmt.Println("oioia")
  decoder   :=  json.NewDecoder(c.Request().Body)
  data      :=  structs.DataMember{}
  err       :=  decoder.Decode(&data)
  checkErr(err)
  fmt.Println(data)

  // data.Status_member = "Tidak Aktif"
  //
  // var password, _ = setRandom.Generate(12, 8, 0, true, true)
  // // fmt.Println(password)
  //
  // var sha = sha1.New()
  // sha.Write([]byte(password))
  // var encrypted = sha.Sum(nil)
  // data.Password = fmt.Sprintf("%x", encrypted)

  // slug      :=  strings.Replace(data.Nama," ","-",-1)
  // isUnique  :=  admin.CheckUniqueSlug(slug)
  //
  // if data.Foto == "" {
  //   data.Foto = "member.png"
  // }
  //
  // if isUnique == true {
  //   data.Slug = slug
  //   // admin.CreateMember(data)
  // } else {
  //   data.Slug = strings.Join([]string{slug,data.Nik},"-")
  //   // admin.CreateMember(data)
  // }
  fmt.Println("oioi")

  file, handler, _ := c.Request().FormFile("file")
  defer file.Close()

  dst, _ := os.OpenFile("assets/images/test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
  defer dst.Close()

  io.Copy(dst, file)

  fmt.Println(handler.Filename)

  // id_lokasi := fmt.Sprintf("%v", data.Id_lokasi)
  // admin.UpdateKamar(id_lokasi, data.No_kamar)
  // fmt.Println(data)

  return c.String(http.StatusOK,"aa")
  // return c.Redirect(http.StatusMovedPermanently, "/")
  // return c.JSON(http.StatusOK, data) // <- blum fix
}

func UpdateMember(c echo.Context) error {
  decoder   :=  json.NewDecoder(c.Request().Body)
  data      :=  structs.DataMember{}
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

func CheckEmailMember(c echo.Context) error {
  email     := c.FormValue("email")
  isValid   := admin.CheckEmailMember(email)

  if isValid == true {
    message := "true"
    return c.String(http.StatusOK, message)
  } else {
    message := "false"
    return c.String(http.StatusOK, message)
  }
}
