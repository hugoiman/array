package admin

import (
  "array/db"
  "array/structs"
)

func GetAdministrasi() structs.Administrasi {
  con       :=  db.Connect()
  // query   :=  "SELECT * FROM administrasi"
  query     :=  "SELECT a.id_administrasi, a.id_member, a.tipe_pembayaran, a.check_in, check_out, a.tgl_pembayaran, a.jumlah_pembayaran, a.total, a.tagihan, a.status, m.nama FROM administrasi a JOIN member m ON a.id_member = m.id_member"
  rows, err :=  con.Query(query)

  checkErr(err)
  defer rows.Close()

  administrasi  :=  structs.Administrasi{}
  data          :=  structs.DataAdministrasi{}

  for rows.Next() {
    err := rows.Scan(
      &data.Id_administrasi, &data.Id_member, &data.Tipe_pembayaran, &data.Check_in, &data.Check_out,
      &data.Tgl_pembayaran, &data.Jumlah_pembayaran, &data.Total, &data.Tagihan, &data.Status,
      &data.Member.Nama,
    )
    data.CustCheck_in   = data.Check_in.Format("02-Jan-2006")
    data.CustCheck_out  = data.Check_out.Format("02-Jan-2006")

    checkErr(err)
    administrasi.Administrasi = append(administrasi.Administrasi, data)
  }
  defer con.Close()

  return administrasi
}
