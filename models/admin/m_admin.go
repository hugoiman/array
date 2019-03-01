package admin

import (
  "array/db"
)

func GetAdmin(id_admin string) DataAdmin {
  con     :=  db.Connect()
  query   :=  "select * from admin where id_admin = ?"

  result  :=  DataAdmin{}
  err     :=  con.QueryRow(query, id_admin).Scan(
    &result.Id_admin, &result.Email,  &result.Nama, &result.Password, &result.Nik,
    &result.No_hp, &result.Alamat, &result.Foto, &result.Tgl_lahir,  &result.Slug,
  )

  checkErr(err)

  defer con.Close()

  return result
}
