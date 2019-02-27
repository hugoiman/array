package auth

import (
  "net/http"
  "fmt"
	"array/models/auth"
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
    return c.Redirect(http.StatusMovedPermanently, "/informasi")
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

    dataSession := auth.DataSession{}

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

func SetSession(c echo.Context, data auth.DataSession) {
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
