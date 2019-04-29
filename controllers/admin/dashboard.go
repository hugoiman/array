package admin

import (
  "fmt"
  "net/http"
  "array/models/admin"
  "array/structs"
  "github.com/labstack/echo"
  "github.com/gorilla/sessions"
  "os"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func ShowDashboard(c echo.Context) error {
  session, _  :=  store.Get(c.Request(), "session")
  slug        :=  fmt.Sprintf("%v", session.Values["slug"])
  data_admin  :=  admin.GetAdmin(slug)

  data := struct {
    Admin     structs.DataAdmin
    Nav       string
  } {
    data_admin,
    "Dashboard",
  }

  // fmt.Printf("%+v\n",data)
  return c.Render(http.StatusOK, "dashboard.html", data)
}
