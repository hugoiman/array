package admin

import (
  "array/db"
  // "fmt"
  "array/structs"
)

var startMonth = "1"
var lastMonth  = "3"
var startYear  = "2019"
var lastYear   = "2019"

func GetPemasukan() structs.PemasukanBulanan {
  con     :=  db.Connect()
  query   :=  "SELECT check_in, SUM(jumlah_pembayaran) FROM administrasi WHERE MONTH(check_in) BETWEEN ? AND ? AND YEAR(check_in) BETWEEN ? AND ? GROUP BY DATE_FORMAT(check_in,'%Y %M')"

  rows, err :=  con.Query(query, startMonth, lastMonth, startYear, lastYear)

  checkErr(err)
  defer rows.Close()

  pemasukan_bulanan := structs.PemasukanBulanan{}
  data              := structs.DataPemasukanBulanan{}

  for rows.Next() {
    err := rows.Scan(&data.WaktuPemasukan, &data.JumlahPemasukan)
    data.BulanPemasukan = data.WaktuPemasukan.Format("2006,01")
    checkErr(err)
    pemasukan_bulanan.PemasukanBulanan = append(pemasukan_bulanan.PemasukanBulanan, data)
  }
  defer con.Close()

  return pemasukan_bulanan
}

func GetPengeluaran() structs.Pengeluaran {
  con     :=  db.Connect()
  x       :=  "2019-02-__"
  query   :=  "SELECT tgl_pengeluaran, jumlah FROM pengeluaran WHERE tgl_pengeluaran LIKE ?"

  rows, err :=  con.Query(query,x)

  checkErr(err)
  defer rows.Close()

  pengeluaran := structs.Pengeluaran{}
  data        := structs.DataPengeluaran{}

  for rows.Next() {
    err := rows.Scan(&data.TglPengeluaran, &data.JumlahPengeluaran)
    checkErr(err)
    pengeluaran.Pengeluaran = append(pengeluaran.Pengeluaran, data)
  }
  defer con.Close()

  return pengeluaran
}
