package member

import (
  "array/db"
  "array/structs"
)

func GetMember(slug string) structs.DataMember {
  con     :=  db.Connect()
  query   :=  "SELECT * FROM member JOIN lokasi_kos ON member.id_lokasi = lokasi_kos.id_lokasi WHERE slug = ?"
  result  :=  structs.DataMember{}
  err     :=  con.QueryRow(query, slug).Scan(
    &result.Id_member, &result.Id_lokasi, &result.Email, &result.Nama, &result.Password, &result.Nik,  &result.No_hp,
    &result.Foto, &result.Tgl_lahir, &result.Pekerjaan, &result.Alamat_asal, &result.No_kamar, &result.Tgl_gabung,
    &result.Status_member, &result.Tipe_pembayaran, &result.Username_wifi, &result.Slug,
    &result.Universitas, &result.Jurusan, &result.Angkatan,
    &result.Nama_kerabat1, &result.Nama_kerabat2, &result.Hubungan1, &result.Hubungan2, &result.No_hp1, &result.No_hp2,
    &result.No_stnk, &result.Perusahaan, &result.Jabatan,
    &result.Lokasi.Id_lokasi, &result.Lokasi.Cabang, &result.Lokasi.Alamat,
  )

  result.CustTgl_lahir   = result.Tgl_lahir.Format("02 January 2006")
  result.CustTgl_gabung  = result.Tgl_gabung.Format("02 January 2006")

  checkErr(err)
  defer con.Close()

  return result
}
