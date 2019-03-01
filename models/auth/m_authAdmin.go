package auth

import (
	"array/db"
)

func AuthAdmin(email, password string) bool {
	var authEmail string
	con     :=  db.Connect()
  query   :=  "SELECT email FROM admin WHERE email = ? AND password = ?"
  err     :=  con.QueryRow(query, email, password).Scan(&authEmail)

	defer con.Close()

	if err == nil {
		return true
	} else {
		return false
	}
}

func GetSessionAdmin(email string) SessionAdmin {
  con     :=  db.Connect()
  query   :=  "SELECT id_admin, nama, slug FROM admin WHERE email = ?"

  result  :=  SessionAdmin{}
  err     :=  con.QueryRow(query, email).Scan(&result.Id_admin, &result.Nama, &result.Slug)

  checkErr(err)
  defer con.Close()

  return result
}
