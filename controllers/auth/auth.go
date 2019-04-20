package auth

import (
  "net/http"
  "fmt"
  c_email "array/controllers/email"
	"array/models/auth"
  "array/models/member"
  "array/models/admin"
  "array/structs"
  "github.com/labstack/echo"
  "crypto/sha1"
  "github.com/gorilla/sessions"
  "os"
  setRandom "github.com/sethvargo/go-password/password"
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
  } else {
    session, _  :=  store.Get(c.Request(), "session")
    slug        :=  fmt.Sprintf("%v", session.Values["slug"])
    if session.Values["id_member"] != nil {
      return c.Redirect(http.StatusMovedPermanently, "/informasi/" + slug)
    } else if session.Values["id_admin"] != nil {
      return c.Redirect(http.StatusMovedPermanently, "/information/status/aktif")
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
  authentic    := auth.Auth(email, encryptedString)

  // fmt.Println("%T\n", result)
  // fmt.Println(authentic)

  if authentic == true {
    result       := auth.GetSession(email)
    SetSession(c, result)
    dataSession := structs.SessionMember{}

    dataSession.Id_member = result.Id_member
    dataSession.Nama = result.Nama
    dataSession.Slug = result.Slug

    // fmt.Printf("%+v\n",dataSession)
    // fmt.Println(reflect.TypeOf(informasi))

    return c.String(http.StatusOK, result.Slug)
    // return c.Redirect(http.StatusMovedPermanently, "/informasi")
	} else {
    message := "false"
    return c.String(http.StatusOK, message)
	}

}

func ResetPassword(c echo.Context) error {
  email :=  c.FormValue("email")
  isAny :=  auth.CheckEmail(email)

  if isAny == true {
    random_password, _  :=  setRandom.Generate(12, 8, 0, true, true)
    to      :=  email
    subject :=  "[Array] Password Reset"
    message :=  "Hello, <b>you</b>, your password is "+random_password

    c_email.SendEmail(to, subject, message)

    return c.String(http.StatusOK, "true")
  } else {
    return c.String(http.StatusOK, "false")
  }
}

func SetSession(c echo.Context, data structs.SessionMember) {
  session, _ := store.Get(c.Request(), "session")
  session.Values["id_member"] = data.Id_member
  session.Values["nama"] = data.Nama
  session.Values["slug"] = data.Slug
  // session.Options = &sessions.Options{   // Atur expired
  //   Path:     "/",
  //   MaxAge:   86400 * 7,
  //   HttpOnly: true,
  // }
  session.Save(c.Request(), c.Response())
  // fmt.Println("session: ", len(session.Values))
}

func CheckSession(c echo.Context) bool {
  session, _ := store.Get(c.Request(), "session")
  if len(session.Values) == 0 {
    return false
  } else {
    return true
  }
}

func Err404(c echo.Context) error {
  session, _ := store.Get(c.Request(), "session")
  slug :=  fmt.Sprintf("%v", session.Values["slug"])

  if session.Values["id_member"] != nil {
    data_member := member.GetMember(slug)
    data := struct {
      Member         structs.DataMember
      Nav            string
    } {
      data_member,
      "Error 404",
    }

    // fmt.Println("eroor member")
    return c.Render(http.StatusOK, "error404.html", data)
  } else if session.Values["id_admin"] != nil {
    data_admin := admin.GetAdmin(slug)
    data := struct {
      Admin         structs.DataAdmin
      Nav           string
    } {
      data_admin,
      "Error 404",
    }
    // fmt.Println("eroor admin")
    return c.Render(http.StatusOK, "err404.html", data)
  } else {
    return c.JSON(http.StatusOK, "err404 boy")
  }
  return nil
}
