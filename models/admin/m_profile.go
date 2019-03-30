package admin

import (
  "array/db"
  "array/structs"
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

func UpdateProfil(data structs.DataAdmin) {
  con     :=  db.Connect()
  query   :=  "UPDATE admin SET nama = ?, email = ?, no_hp = ?, nik = ?, tgl_lahir = ?, alamat = ? WHERE id_admin = ?"
  _, err  :=  con.Exec(query, data.Nama, data.Email, data.No_hp, data.Nik, data.CustTgl_lahir, data.Alamat, data.Id_admin)

  checkErr(err)
  defer con.Close()
}

func UpdateFoto(id_admin, new_foto string) {
  con     :=  db.Connect()
  query   :=  "UPDATE admin SET foto = ? WHERE id_admin = ?"
  _, err  :=  con.Exec(query, new_foto, id_admin)

  checkErr(err)
  defer con.Close()
}
