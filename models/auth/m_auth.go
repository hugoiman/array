package auth

import (
	// "fmt"
	"array/db"
	"array/structs"
)

func checkErr(err error){
  if err != nil {
    panic(err.Error())
  }
}

func Auth(email, password string) bool {
	var authEmail string
	con     :=  db.Connect()
  query   :=  "SELECT email FROM member WHERE email = ? AND password = ?"
  err     :=  con.QueryRow(query, email, password).Scan(&authEmail)

	defer con.Close()

	if err == nil {
		return true
	} else {
		return false
	}
}

func CheckEmail(email string) bool {
	var isAny string
	con     :=  db.Connect()
  query   :=  "SELECT email FROM member WHERE email = ?"
  err     :=  con.QueryRow(query, email).Scan(&isAny)

	defer con.Close()

	if err == nil {
		return true
	} else {
		return false
	}
}

func GetSession(email string) structs.SessionMember {
  con     :=  db.Connect()
  query   :=  "SELECT id_member, nama, slug FROM member WHERE email = ?"

  result  :=  structs.SessionMember{}
  err     :=  con.QueryRow(query, email).Scan(&result.Id_member, &result.Nama, &result.Slug)

  checkErr(err)
  defer con.Close()

  return result
}
