package member

import (
  "array/db"
)

func GetAdministrasi(id_member string) Administrasi {
  con         :=  db.Connect()
  query       :=  "SELECT * FROM administrasi where id_member = ?"
  rows, err   :=  con.Query(query, id_member)

  checkErr(err)

  data    :=  DataAdministrasi{}
  info    :=  Administrasi{}

  for rows.Next() {
    err   :=  rows.Scan(
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

func GetDetailTagihan(id_member string) InfoTagihan {
  con       :=  db.Connect()
  query     :=  "SELECT check_in, check_out, tagihan FROM administrasi WHERE id_member = ? AND status = 'Belum Lunas'"
  rows, err :=  con.Query(query, id_member)

  checkErr(err)

  data    :=  DataTagihan{}
  info    :=  InfoTagihan{}

  for rows.Next() {
    err   :=  rows.Scan(&data.Check_in, &data.Check_out, &data.Tagihan,)
    checkErr(err)
    info.InfoTagihan = append(info.InfoTagihan, data)
  }

  checkErr(err)
  defer con.Close()

  return info
}

func GetTagihan(id_member string) (string, int) {
  var status string
  var tagihan int
  con     :=  db.Connect()
  query   :=  "SELECT(SELECT SUM(tagihan) FROM administrasi WHERE id_member = ?) as tagihan,(SELECT status FROM administrasi WHERE id_member = ? ORDER BY id_administrasi DESC LIMIT 1) as status"
  err     :=  con.QueryRow(query, id_member, id_member).Scan(&tagihan, &status)

  checkErr(err)

  defer con.Close()

  return status, tagihan
}
