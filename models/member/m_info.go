package member

import (
  "array/db"
)

func GetInfo() Informasi{
  con     :=  db.Connect()
  query   :=  "select id_informasi, judul, keterangan, foto, waktu, status from informasi"
  rows, err := con.Query(query)

  checkErr(err)
  defer rows.Close()

  info := Informasi{}
  data := DataInformasi{}

  for rows.Next() {
    err := rows.Scan(&data.Id_informasi, &data.Judul, &data.Keterangan, &data.Foto, &data.Waktu, &data.Status)
    checkErr(err)
    info.Informasi = append(info.Informasi, data)
  }
  defer con.Close()

  return info
}
