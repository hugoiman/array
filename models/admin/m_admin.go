package admin

import (
  "array/db"
)

func GetAdmin(slug string) DataAdmin {
  con     :=  db.Connect()
  query   :=  "select * from admin where slug = ?"

  result  :=  DataAdmin{}
  err     :=  con.QueryRow(query, slug).Scan(
    &result.Id_admin, &result.Email,  &result.Nama, &result.Password, &result.Nik,
    &result.Foto, &result.No_hp, &result.Alamat, &result.Tgl_lahir,  &result.Slug,
  )

  checkErr(err)
  defer con.Close()

  return result
}