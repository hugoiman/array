package auth

import (
  "net/http"
  "fmt"
	"array/models/auth"
  "array/models/member"
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
  var data map[string]interface{}
  data = map[string]interface{}{
    "name":      "tes",
    "grade":     2,
  }

  session := CheckSession(c)

  if session == false {
    result := auth.Index()
    return c.Render(http.StatusOK, "index.html", result)
  }else {
    informasi := member.GetInfo()
    fmt.Printf("%+v\n",informasi)
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
    SetSession(c, result)

    dataSession := auth.DataMember{}

    dataSession.Id_member = result.Id_member
    dataSession.Nama = result.Nama
    dataSession.Email = result.Email
    dataSession.Password = result.Password

    // session.Member = append(session.Member,dataSession)
    fmt.Printf("%+v\n",dataSession)
    // fmt.Println(reflect.TypeOf(informasi))

    // return c.Render(http.StatusOK, "home.html", informasi)
    return c.Redirect(http.StatusMovedPermanently, "/informasi")
	} else {
    fmt.Println("Login gagal")
		//login failed
		// return c.Redirect(http.StatusTemporaryRedirect, "/")
    return c.Redirect(http.StatusMovedPermanently, "/")
	}

}

func SetSession(c echo.Context, data auth.DataMember) {
  session, _ := store.Get(c.Request(), "session")
  session.Values["id_member"] = data.Id_member
  session.Values["nama"] = data.Nama
  session.Values["email"] = data.Email
  session.Values["password"] = data.Password
  session.Save(c.Request(), c.Response())

  fmt.Println("session: ", len(session.Values))
  // fmt.Println(session)
}

func CheckSession(c echo.Context) bool {
  session, _ := store.Get(c.Request(), "session")
  if len(session.Values) == 0 {
    return false
  }else {
    return true
  }
}
