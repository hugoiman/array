package admin

import (
  "array/db"
  "array/structs"
)

func GetAllLokasi() structs.Lokasi{
  con       :=  db.Connect()
  query     :=  "SELECT * FROM lokasi_kos"
  rows, err :=  con.Query(query)

  checkErr(err)
  defer rows.Close()

  lokasi    :=  structs.Lokasi{}
  data      :=  structs.DataLokasi{}

  for rows.Next() {
    err := rows.Scan(
      &data.Id_lokasi, &data.Cabang, &data.Alamat,
    )

    checkErr(err)
    lokasi.Lokasi = append(lokasi.Lokasi, data)
  }
  defer con.Close()

  return lokasi
}

func CreateLokasi(cabang, alamat string) {
	con     :=  db.Connect()
	_, err 	:=  con.Exec("INSERT INTO lokasi_kos (cabang, alamat) VALUES (?,?)", cabang, alamat)

	checkErr(err)
	defer con.Close()
}

func UpdateLokasi(id_lokasi, cabang, alamat string) {
  con     :=  db.Connect()
  query   :=  "UPDATE lokasi_kos SET cabang = ?, alamat = ? WHERE id_lokasi = ?"
  _, err  :=  con.Exec(query, cabang, alamat, id_lokasi)

  checkErr(err)
  defer con.Close()
}

func DeleteLokasi(id_lokasi string) {
	con     :=  db.Connect()
	_, err 	:=  con.Exec("DELETE FROM lokasi_kos WHERE id_lokasi = ?", id_lokasi)

	checkErr(err)
	defer con.Close()
}
