package admin

import (
  "array/db"
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

func CheckEmail(new_email string) bool {
  var result_email string
  con     :=  db.Connect()
  query   :=  "select email from admin where email = ?"
  err     :=  con.QueryRow(query, new_email).Scan(&result_email)

  defer con.Close()

  if err == nil {
    return false
  } else {
    return true
  }
}

func UpdateBiodata(id_admin int, nama string) {
  con     :=  db.Connect()
  query   :=  "UPDATE admin SET nama = ? WHERE id_admin = ?"
  _, err  :=  con.Exec(query, nama, id_admin)

  checkErr(err)
  defer con.Close()
}
