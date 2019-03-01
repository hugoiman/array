package admin

import (
  "array/db"
  "fmt"
)

func GetPassword(id_admin string) string {
  var password string
  con     :=  db.Connect()
  query   :=  "select password from admin where id_admin = ?"
  err     :=  con.QueryRow(query, id_admin).Scan(&password)

  checkErr(err)
  defer con.Close()

  return password
}

func UpdatePassword(id_admin, password string) {
  con     :=  db.Connect()
  query   :=  "UPDATE admin SET password = ? WHERE id_admin = ?"
  _, err  :=  con.Exec(query, password, id_admin)
  
  checkErr(err)
  defer con.Close()
}
