package admin

import (
  "net/http"
  "array/models/admin"
  "github.com/labstack/echo"
  "fmt"
  "github.com/gorilla/sessions"
  "os"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func ShowDashboard(c echo.Context) error{
  session, _ := store.Get(c.Request(), "session")
  id_admin   := fmt.Sprintf("%v", session.Values["id_admin"])
  dataAdmin  := admin.GetAdmin(id_admin)

  data := struct {
    Admin     admin.DataAdmin
    Nav       string
  } {
    dataAdmin,
    "informasi",
  }

  // fmt.Printf("%+v\n",data)
  return c.Render(http.StatusOK, "dashboard.html", data)
}
