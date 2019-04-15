package structs

import(
  "time"
)

type PemasukanBulanan struct {
  PemasukanBulanan          []DataPemasukanBulanan   `json:"pemasukan_bulanan"`
}
type DataPemasukanBulanan struct {
  WaktuPemasukan            time.Time         `json:"waktu_pemasukan"`
  JumlahPemasukan           int               `json:"jumlah_pemasukan"`

  BulanPemasukan            string
}

type Pengeluaran struct {
  Pengeluaran               []DataPengeluaran `json:"pengeluaran"`
}
type DataPengeluaran struct {
  TglPengeluaran            time.Time         `json:"tgl_pengeluaran"`
  JumlahPengeluaran         int               `json:"jumlah_pengeluaran"`

  CustTglPengeluaran        string
}
