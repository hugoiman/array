package member

import (
  "array/db"
  "fmt"
)

func GetPassword(id_member string) string {
  var password string
  con     :=  db.Connect()
  query   :=  "select password from member where id_member = ?"
  err     :=  con.QueryRow(query, id_member).Scan(&password)

  checkErr(err)
  defer con.Close()

  return password
}

func UpdatePassword(id_member, password string) {
  con     :=  db.Connect()
  query   :=  "UPDATE member SET password = ? WHERE id_member = ?"
  _, err  :=  con.Exec(query, password, id_member)
  
  checkErr(err)

  defer con.Close()
}
