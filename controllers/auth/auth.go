package auth

import (
  "net/http"
  "fmt"
	"array/models/auth"
  "array/models/member"
  "array/models/admin"
  "github.com/labstack/echo"
  "crypto/sha1"
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

func Index(c echo.Context) error{
  session := CheckSession(c)
  data := struct {
    o string
  }{
    "1",
  }

  if session == false {
    return c.Render(http.StatusOK, "index.html", data)
  }else {
    session, _ := store.Get(c.Request(), "session")
    if session.Values["id_member"] != "" {
      return c.Redirect(http.StatusMovedPermanently, "/informasi")
    } else {
      return c.Redirect(http.StatusMovedPermanently, "/dashboard")
    }

  }
}

func Login(c echo.Context) error{
	email        := c.FormValue("email")
  password     := c.FormValue("password")

  var sha = sha1.New()
  sha.Write([]byte(password))
  var encrypted = sha.Sum(nil)
  var encryptedString = fmt.Sprintf("%x", encrypted)

  authentic       := auth.Auth(email, encryptedString)

  // fmt.Println("%T\n", result)
  fmt.Println(authentic)

  if authentic == true {
    result       := auth.GetSession(email)
    SetSession(c, result)

    dataSession := auth.SessionMember{}

    dataSession.Id_member = result.Id_member
    dataSession.Nama = result.Nama
    dataSession.Slug = result.Slug

    // fmt.Printf("%+v\n",dataSession)
    // fmt.Println(reflect.TypeOf(informasi))
    message := "true"
    return c.String(http.StatusOK, message)
    // return c.Redirect(http.StatusMovedPermanently, "/informasi")
	} else {
    message := "false"
    return c.String(http.StatusOK, message)
	}

}

func SetSession(c echo.Context, data auth.SessionMember) {
  session, _ := store.Get(c.Request(), "session")
  session.Values["id_member"] = data.Id_member
  session.Values["nama"] = data.Nama
  session.Values["slug"] = data.Slug
  session.Save(c.Request(), c.Response())

  // fmt.Println("session: ", len(session.Values))
}

func CheckSession(c echo.Context) bool {
  session, _ := store.Get(c.Request(), "session")
  if len(session.Values) == 0 {
    return false
  }else {
    return true
  }
}

func Err404(c echo.Context) error {
  session, _ := store.Get(c.Request(), "session")
  if session.Values["id_member"] != nil {
    id_member :=  fmt.Sprintf("%v", session.Values["id_member"])
    data_member := member.GetMember(id_member)
    data := struct {
      Member         member.DataMember
      Nav            string
    } {
      data_member,
      "no nav",
    }
    return c.Render(http.StatusOK, "error404.html", data)

  } else if session.Values["id_admin"] != nil {
    id_admin :=  fmt.Sprintf("%v", session.Values["id_admin"])
    data_admin := admin.GetAdmin(id_admin)
    data := struct {
      Admin         admin.DataAdmin
      Nav           string
    } {
      data_admin,
      "no nav",
    }
    return c.Render(http.StatusOK, "err404.html", data)

  } else {
    fmt.Println("auu")
    return c.JSON(http.StatusOK, "err404 boy")
  }
  return nil
}
