package auth

import (
  "net/http"
  "fmt"
	"array/models/auth"
  "github.com/labstack/echo"
  "crypto/sha1"
  "github.com/gorilla/sessions"
  "os"
)

func checkErr(err error){
  if err != nil {
    panic(err.Error())
  }
}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func Index(c echo.Context) error{
  var data map[string]interface{}
  data = map[string]interface{}{
    "name":      "tes",
    "grade":     2,
  }
  session, _ := store.Get(c.Request(), "session")
  //
  fmt.Println("session: ", len(session.Values))
  // fmt.Println(session.Values["email"])

  if len(session.Values) == 0 {
    result := auth.Index()
    return c.Render(http.StatusOK, "index.html", result)
  }else {
    return c.Render(http.StatusOK, "home.html", data)
  }
}

func Login(c echo.Context) error{
	email := c.FormValue("email")
  password := c.FormValue("password")
  result := auth.Login(email, password)

  var sha = sha1.New()
  sha.Write([]byte(password))
  var encrypted = sha.Sum(nil)

  var encryptedString = fmt.Sprintf("%x", encrypted)
  // fmt.Println("%T\n", result)

  if encryptedString == result.Password {
    fmt.Println("Login sukses")
    dataSession := SetSession(c, result)
    // fmt.Printf("%+v\n",dataSession)
    return c.Render(http.StatusOK, "home.html", dataSession)
	} else {
    fmt.Println("Login gagal")
		//login failed
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

}

func SetSession(c echo.Context, data auth.User) *sessions.Session {
  session, _ := store.Get(c.Request(), "session")
  session.Values["id_user"] = data.Id_user
  session.Values["nama"] = data.Nama
  session.Values["email"] = data.Email
  session.Values["password"] = data.Password
  session.Save(c.Request(), c.Response())

  fmt.Println(session.Values["email"])
  fmt.Println("session: ", len(session.Values))

  return session
}

func Logout(c echo.Context) error{
  session, _ := store.Get(c.Request(), "session")
  session.Options.MaxAge = -1
  session.Save(c.Request(), c.Response())

  return c.Redirect(http.StatusTemporaryRedirect, "/")
}
