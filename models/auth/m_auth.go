package auth

import (
	"fmt"
	"array/db"
)

type User struct {
	Id_user    int     `json:"id_user"`
	Email      string  `json:"email"`
	Nama       string  `json:"nama"`
  Password   string  `json:"password"`
}

type Users struct {
  Users    []User  `json:"user"`
}

func checkErr(err error){
  if err != nil {
    fmt.Println("error")
  }
}

func Index() Users{
  con     :=  db.Connect()
  query   :=  "select id_user, nama, email, password from user"

  rows, err := con.Query(query)

  checkErr(err)

  defer rows.Close()

  result := Users{}
  user := User{}

  for rows.Next() {
    err := rows.Scan(&user.Id_user, &user.Nama, &user.Email, &user.Password)
    checkErr(err)
    result.Users = append(result.Users, user)
    // fmt.Println(result.Name)
    // fmt.Println(result.Email)
  }
  defer con.Close()

  return result
}

func Login(email, password string) User {
  con     :=  db.Connect()
  query   :=  "select id_user, email, nama, password from user where email = ?"

  result  :=  User{}
  err     :=  con.QueryRow(query, email).Scan(&result.Id_user, &result.Email,  &result.Nama, &result.Password)

  checkErr(err)

  defer con.Close()

  return result
}
