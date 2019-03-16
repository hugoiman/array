package admin

import (
  "array/db"
)

func GetAllInformasi(status string) Informasi {
  con     :=  db.Connect()
  if status == "semua" {
    status = "%tif"
  }
  query   :=  "SELECT id_informasi, judul, keterangan, foto, waktu, status FROM informasi WHERE status LIKE ? ORDER BY id_informasi DESC"
  rows, err := con.Query(query, status)

  checkErr(err)
  defer rows.Close()

  informasi := Informasi{}
  data := DataInformasi{}

  for rows.Next() {
    err := rows.Scan(&data.Id_informasi, &data.Judul, &data.Keterangan, &data.Foto, &data.Waktu, &data.Status)
    data.CustomTime = data.Waktu.Format("02 Jan 2006")
    checkErr(err)
    informasi.Informasi = append(informasi.Informasi, data)
  }
  defer con.Close()

  return informasi
}

func GetInformasi(id_informasi string) DataInformasi {
  con     :=  db.Connect()
  query   :=  "SELECT * FROM informasi WHERE id_informasi = ?"

  data    :=  DataInformasi{}
  err     :=  con.QueryRow(query, id_informasi).Scan(&data.Id_informasi, &data.Judul, &data.Keterangan, &data.Foto, &data.Waktu, &data.Status)

  checkErr(err)
  defer con.Close()

  return data
}

func CreateInformasi(judul, keterangan, waktu, status string) {
	con     :=  db.Connect()
	_, err 	:=  con.Exec("INSERT INTO informasi (judul, keterangan, waktu, status) VALUES (?, ?, ?, ?)", judul, keterangan, waktu, status)

	checkErr(err)
	defer con.Close()
}

func DeleteInformasi(id_informasi string) {
	con     :=  db.Connect()
	_, err 	:=  con.Exec("DELETE FROM informasi WHERE id_informasi = ?", id_informasi)

	checkErr(err)
	defer con.Close()
}

// func UpdateInformasi(id_member int, data struct) {
// 	con     :=  db.Connect()
// 	_, err 	:=  con.Exec("UPDATE member SET ... WHERE id_member = ?", id_member)
//
// 	checkErr(err)
// 	defer con.Close()
// }

func CountInformasi() (string, string, string, string) {
  var semua, aktif, non, draft string
  con     :=  db.Connect()
  query   :=  "SELECT (SELECT COUNT(status) FROM informasi) as semua, (SELECT COUNT(status) FROM informasi WHERE status = 'aktif') as aktif, (SELECT COUNT(status) FROM informasi WHERE status = 'non-aktif') as non, (SELECT COUNT(status) FROM informasi WHERE status = 'draft') as draft"
  err     :=  con.QueryRow(query).Scan(&semua, &aktif, &non, &draft)

  checkErr(err)
  defer con.Close()

  return semua, aktif, non, draft
}
