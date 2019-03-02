package member

import(
  "github.com/labstack/echo"
  "fmt"
)

func checkErr(err error){
  if err != nil {
    panic(err.Error())
  }
}

func checkSlug(c echo.Context, slug string) error {
  session, _  := store.Get(c.Request(), "session")
  sessSlug := fmt.Sprintf("%v", session.Values["slug"])
  if slug ==  sessSlug {
    return nil
  } else {
    return echo.NotFoundHandler(c)
  }
}
