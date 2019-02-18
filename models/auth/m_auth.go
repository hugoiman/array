package auth

import (
	// "fmt"
	"array/db"
)

func checkErr(err error){
  if err != nil {
    panic(err.Error())
  }
}

func Login(email, password string) DataMember {
  con     :=  db.Connect()
  query   :=  "select id_member, nama, password, slug from member where email = ?"

  result  :=  DataMember{}
  err     :=  con.QueryRow(query, email).Scan(&result.Id_member,  &result.Nama, &result.Password, &result.Slug)

  checkErr(err)

  defer con.Close()

  return result
}

//
// func Index() Member{
//   con     :=  db.Connect()
//   query   :=  "select id_member, nama, email, password, slug from member"
//
//   rows, err := con.Query(query)
//
//   checkErr(err)
//
//   defer rows.Close()
//
//   result := Member{}
//   member := DataMember{}
//
//   for rows.Next() {
//     err := rows.Scan(&member.Id_member, &member.Nama, &member.Email, &member.Password, &member.Slug)
//     checkErr(err)
//     result.Member = append(result.Member, member)
//     // fmt.Println(result.Name)
//   }
//   defer con.Close()
//
//   return result
// }
