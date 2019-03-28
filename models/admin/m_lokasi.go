package admin

import (
  "array/db"
  "array/structs"
)

func GetLokasi() structs.Lokasi{
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
