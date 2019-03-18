package structs

import(
  "time"
)

type Informasi struct {
  Informasi             []DataInformasi     `json:"informasi"`
}
type DataInformasi struct {
  Id_informasi          int            `json:"id_informasi"`
  Judul                 string         `json:"judul"`
  Keterangan            string         `json:"keterangan"`
  Foto                  string         `json:"foto"`
  Waktu                 time.Time      `json:"waktu"`
  Status                string         `json:"status"`

  CustWaktu             string
}
