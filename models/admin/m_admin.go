package admin

import (
  "array/db"
  "array/structs"
)

func GetAdmin(slug string) structs.DataAdmin {
  con     :=  db.Connect()
  query   :=  "select * from admin where slug = ?"

  result  :=  structs.DataAdmin{}
  err     :=  con.QueryRow(query, slug).Scan(
    &result.Id_admin, &result.Email,  &result.Nama, &result.Password, &result.Nik,
    &result.Foto, &result.No_hp, &result.Alamat, &result.Tgl_lahir,  &result.Slug,
  )

  result.CustTgl_lahir = result.Tgl_lahir.Format("02 January 2006")

  checkErr(err)
  defer con.Close()

  return result
}
