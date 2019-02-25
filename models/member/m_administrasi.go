package member

import (
  "array/db"
)

func GetAdministrasi(id_member string) Administrasi {
  con     :=  db.Connect()
  query   :=  "select * from administrasi where id_member = ?"

  rows, err := con.Query(query, id_member)

  checkErr(err)

  data    :=  DataAdministrasi{}
  info    :=  Administrasi{}

  for rows.Next() {
    err     :=  rows.Scan(
      &data.Id_administrasi, &data.Id_member,
      &data.Tipe_pembayaran, &data.Tgl_pembayaran,
      &data.Check_in, &data.Check_out,
      &data.Jumlah_pembayaran, &data.Total,
      &data.Tagihan, &data.Status,
    )
    checkErr(err)
    info.Administrasi = append(info.Administrasi, data)
  }

  defer con.Close()

  return info
}
