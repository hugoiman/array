package admin

import (
  "fmt"
  "net/http"
  "encoding/json"
  // "array/models/admin"
  "github.com/labstack/echo"
  "io/ioutil"
  "os"
)

// func UpdateKamar(c echo.Context, id_lokasi, no_kamar string) error {
//   admin.UpdateKamar(id_lokasi, no_kamar)
// }

func GetKamarJson(c echo.Context) error  {
  var path = "api/kamar.json"
  jsonFile, err := os.Open(path)
  // if we os.Open returns an error then handle it
  if err != nil {
  	fmt.Println(err)
  }
  fmt.Println("Successfully Opened users.json")
  // defer the closing of our jsonFile so that we can parse it later on
  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

  return c.JSON(http.StatusOK, result)
}
