package admin

import (
	"array/db"
	"array/structs"
	// "fmt"
)

func GetMembers() structs.Member {
  con     :=  db.Connect()
  query   :=  "SELECT * FROM member JOIN lokasi_kos ON member.id_lokasi = lokasi_kos.id_lokasi"
  rows, err := con.Query(query)

  checkErr(err)
  defer rows.Close()

  member := structs.Member{}
  data := structs.DataMember{}

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
    data.CustTgl_lahir   = data.Tgl_lahir.Format("02 January 2006")
    data.CustTgl_gabung  = data.Tgl_gabung.Format("02 January 2006")

    checkErr(err)
    member.Member = append(member.Member, data)
  }
  defer con.Close()

  return member
}

func CreateMember(data structs.DataMember) {
	con     :=  db.Connect()
	_, err 	:=  con.Exec("INSERT INTO member (id_lokasi,email,nama,password,nik,no_hp,foto,tgl_lahir,pekerjaan,alamat_asal,no_kamar,status_member,tipe_pembayaran,username_wifi,slug,universitas,jurusan,angkatan,nama_kerabat1,nama_kerabat2,hubungan1,hubungan2,no_hp1,no_hp2,no_stnk,perusahaan,jabatan) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		data.Id_lokasi, data.Email, data.Nama, data.Password, data.Nik, data.No_hp, data.Foto, data.CustTgl_lahir,
		data.Pekerjaan, data.Alamat_asal, data.No_kamar, data.Status_member, data.Tipe_pembayaran, data.Username_wifi,
		data.Slug, data.Universitas, data.Jurusan, data.Angkatan, data.Nama_kerabat1, data.Nama_kerabat2,
		data.Hubungan1, data.Hubungan2, data.No_hp1, data.No_hp2, data.No_stnk, data.Perusahaan, data.Jabatan)
	// fmt.Println(data.Email)
	checkErr(err)
	defer con.Close()
}

// func UpdateMember(id_member int, data struct) {
// 	con     :=  db.Connect()
// 	_, err 	:=  con.Exec("UPDATE member SET ... WHERE id_member = ?", id_member)
//
// 	checkErr(err)
// 	defer con.Close()
// }

func DeleteMember(id_member int) {
	con     :=  db.Connect()
	_, err 	:=  con.Exec("Delete FROM member WHERE id_member = ?", id_member)

	checkErr(err)
	defer con.Close()
}

func CheckUniqueSlug(slug string) bool {
	var result_slug string
  con     :=  db.Connect()
  query   :=  "SELECT slug FROM member WHERE slug = ?"
  err     :=  con.QueryRow(query, slug).Scan(&result_slug)

  defer con.Close()

  if err == nil {
    return false
  } else {
    return true
  }
}

func CheckEmailMember(email string) bool {
  var result_email string
  con     :=  db.Connect()
  query   :=  "select email from member where email = ?"
  err     :=  con.QueryRow(query, email).Scan(&result_email)

  defer con.Close()

  if err == nil {
    return false
  } else {
    return true
  }
}
