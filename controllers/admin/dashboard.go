package admin

import (
  "net/http"
  "array/models/admin"
  "github.com/labstack/echo"
  "github.com/gorilla/sessions"
  "os"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func ShowDashboard(c echo.Context) error {
  slug := c.Param("slug")
  checkSlug(c, slug)
  data_admin  := admin.GetAdmin(slug)

  data := struct {
    Admin     admin.DataAdmin
    Nav       string
  } {
    data_admin,
    "Dashboard",
  }

  // fmt.Printf("%+v\n",data)
  return c.Render(http.StatusOK, "dashboard.html", data)
}
