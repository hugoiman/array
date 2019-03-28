package admin

import (
	"array/db"
  "array/structs"
)

func GetKamar() structs.Kamar{
  con       :=  db.Connect()
  query     :=  "SELECT * FROM kamar"
  rows, err :=  con.Query(query)

  checkErr(err)
  defer rows.Close()

  kamar    :=  structs.Kamar{}
  data      :=  structs.DataKamar{}

  for rows.Next() {
    err := rows.Scan(
      &data.Id_kamar, &data.Id_lokasi, &data.No_kamar, &data.Status_kamar,
    )

    checkErr(err)
    kamar.Kamar = append(kamar.Kamar, data)
  }
  defer con.Close()

  return kamar
}

func UpdateKamar(id_lokasi, no_kamar string)  {
	con     :=  db.Connect()
  query   :=  "UPDATE kamar SET status_kamar = 'Terisi' WHERE id_lokasi = ? && no_kamar = ?"
  _, err  :=  con.Exec(query, id_lokasi, no_kamar)

  checkErr(err)
  defer con.Close()
}
