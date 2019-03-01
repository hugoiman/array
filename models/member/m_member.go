package member

import (
  "array/db"
)

func GetMember(id_member string) DataMember {
  con     :=  db.Connect()
  query   :=  "select * from member where id_member = ?"
  result  :=  DataMember{}
  err     :=  con.QueryRow(query, id_member).Scan(
    &result.Id_member, &result.Id_lokasi, &result.Email,  &result.Nama, &result.Password, &result.Nik,  &result.No_hp,
    &result.Foto, &result.Tgl_lahir,  &result.Pekerjaan, &result.Alamat_asal, &result.No_kamar,  &result.Tgl_gabung,
    &result.Status_member, &result.Tipe_pembayaran,  &result.Username_wifi, &result.Slug,
    &result.Universitas,  &result.Jurusan, &result.Angkatan,
    &result.Nama_kerabat1,  &result.Nama_kerabat2, &result.Hubungan1, &result.Hubungan2,  &result.No_hp1, &result.No_hp2,
    &result.No_stnk,  &result.Perusahaan, &result.Jabatan,
  )

  checkErr(err)
  defer con.Close()

  return result
}
