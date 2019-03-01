package admin

import(
  "time"
)

type DataAdmin struct {
  Id_admin              int            `json:"id_admin"`
  Nama                  string         `json:"nama"`
  Email                 string         `json:"email"`
  Password              string         `json:"password"`
  Nik                   string         `json:"nik"`
  No_hp                 string         `json:"no_hp"`
  Alamat                string         `json:"alamat"`
  Foto                  string         `json:"foto"`
  Tgl_lahir             time.Time      `json:"tgl_lahir"`
  Slug                  string         `json:"slug"`
}

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
}

type Lokasi struct {
  Lokasi                []DataLokasi   `json:"lokasi"`
}
type DataLokasi struct {
  Id_lokasi             int            `json:"id_lokasi"`
  Cabang                string         `json:"cabang"`
  Alamat                string         `json:"alamat"`
}

type Administrasi struct {
  Administrasi          []DataAdministrasi     `json:"administrasi"`
}
type DataAdministrasi struct {
  Id_administrasi       int         `json:"id_administrasi"`
  Id_member             int         `json:"id_member"`
  Tipe_pembayaran       string      `json:"tipe_pembayaran"`
  Check_in              time.Time   `json:"check_in"`
  Check_out             time.Time   `json:"check_out"`
  Tgl_pembayaran        time.Time   `json:"tgl_pembayaran"`
  Jumlah_pembayaran     int         `json:"jumlah_pembayaran"`
  Total                 int         `json:"total"`
  Tagihan               int         `json:"tagihan"`
  Status                string      `json:"status"`
}
