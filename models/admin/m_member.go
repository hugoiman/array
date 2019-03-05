package admin

import (
  // "fmt"
	"array/db"
)

func GetMembers() Member {
  con     :=  db.Connect()
  query   :=  "SELECT * FROM member JOIN lokasi_kos ON member.id_lokasi = lokasi_kos.id_lokasi"
  rows, err := con.Query(query)

  checkErr(err)
  defer rows.Close()

  member := Member{}
  data := DataMember{}

  for rows.Next() {
    err := rows.Scan(
      &data.Id_member, &data.Id_lokasi, &data.Email,  &data.Nama, &data.Password, &data.Nik,  &data.No_hp,
      &data.Foto, &data.Tgl_lahir,  &data.Pekerjaan, &data.Alamat_asal, &data.No_kamar, &data.Tgl_gabung,
      &data.Status_member, &data.Tipe_pembayaran, &data.Username_wifi, &data.Slug,
      &data.Universitas, &data.Jurusan, &data.Angkatan,
      &data.Nama_kerabat1, &data.Nama_kerabat2, &data.Hubungan1, &data.Hubungan2, &data.No_hp1, &data.No_hp2,
      &data.No_stnk, &data.Perusahaan, &data.Jabatan,
      &data.Lokasi.Id_lokasi, &data.Lokasi.Cabang, &data.Lokasi.Alamat,
    )

    checkErr(err)
    member.Member = append(member.Member, data)
  }
  defer con.Close()

  return member
}
