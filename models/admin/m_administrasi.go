package admin

import (
  "array/db"
  "array/structs"
)

func GetAdministrasi() structs.Member {
  con       :=  db.Connect()
  query     :=  "SELECT m.nama, m.id_member, m.id_lokasi, l.id_lokasi, a.id_administrasi, a.id_member, a.tipe_pembayaran, a.check_in, a.check_out, a.tgl_pembayaran, a.jumlah_pembayaran, a.total, a.tagihan, a.status FROM member m JOIN administrasi a ON m.id_member = a.id_member JOIN lokasi_kos l ON l.id_lokasi = m.id_lokasi"
  rows, err :=  con.Query(query)

  checkErr(err)
  defer rows.Close()

  member    :=  structs.Member{}
  data      :=  structs.DataMember{}

  for rows.Next() {
    err := rows.Scan(
      &data.Nama, &data.Id_member, &data.Id_lokasi,
      &data.Lokasi.Id_lokasi,
      &data.Administrasi.Id_administrasi, &data.Administrasi.Id_member, &data.Administrasi.Tipe_pembayaran, &data.Administrasi.Check_in, &data.Administrasi.Check_out,
      &data.Administrasi.Tgl_pembayaran, &data.Administrasi.Jumlah_pembayaran, &data.Administrasi.Total, &data.Administrasi.Tagihan, &data.Administrasi.Status,
    )
    data.Administrasi.CustCheck_in   = data.Administrasi.Check_in.Format("02 Jan 2006")
    data.Administrasi.CustCheck_out  = data.Administrasi.Check_out.Format("02 Jan 2006")

    checkErr(err)
    member.Member = append(member.Member, data)
  }
  defer con.Close()

  return member
}
