package member

import (
  "array/db"
)

func GetMember(id_member string) DataMember {
  con     :=  db.Connect()
  query   :=  "select id_member, email, nama, password from member where id_member = ?"

  result  :=  DataMember{}
  err     :=  con.QueryRow(query, id_member).Scan(&result.Id_member, &result.Email,  &result.Nama, &result.Password)

  checkErr(err)

  defer con.Close()

  return result
}
