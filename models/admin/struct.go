package admin

import(
  "time"
)

type Member struct {
  Member                []DataMember   `json:"member"`
}
type DataMember struct {
  Id_member             int            `json:"id_member"`
  Id_lokasi             int            `json:"id_lokasi"`
  Nama                  string         `json:"nama"`
  Email                 string         `json:"email"`
  Password              string         `json:"password"`
  Nik                   string         `json:"nik"`
  No_hp                 string         `json:"no_hp"`
  Foto                  string         `json:"foto"`
  Tgl_lahir             time.Time      `json:"tgl_lahir"`
  Pekerjaan             string         `json:"pekerjaan"`
  Alamat_asal           string         `json:"alamat_asal"`
  No_kamar              string         `json:"no_kamar"`
  Tgl_gabung            time.Time      `json:"tgl_gabung"`
  Status_member         string         `json:"status_member"`
  Tipe_pembayaran       string         `json:"tipe_pembayaran"`
  Username_wifi         string         `json:"username_wifi"`
  Slug                  string         `json:"slug"`
  Universitas           string         `json:"universitas"`
  Jurusan               string         `json:"jurusan"`
  Angkatan              string         `json:"angkatan"`
  Nama_kerabat1         string         `json:"nama_kerabat1"`
  Nama_kerabat2         string         `json:"nama_kerabat2"`
  Hubungan1             string         `json:"hubungan1"`
  Hubungan2             string         `json:"hubungan2"`
  No_hp1                string         `json:"no_hp1"`
  No_hp2                string         `json:"no_hp2"`
  No_stnk               string         `json:"no_stnk"`
  Perusahaan            string         `json:"perusahaan"`
  Jabatan               string         `json:"jabatan"`

  Lokasi                DataLokasi     `json:"lokasi"`
}

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

type Biodata struct {
  Id_admin              int            `json:"id_admin"`
  Nama                  string         `json:"nama"`
  Email                 string         `json:"email"`
  Password              string         `json:"password"`
  Nik                   string         `json:"nik"`
  No_hp                 string         `json:"no_hp"`
  Alamat                string         `json:"alamat"`
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
  CustomTime            string
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

  Member                DataMember  `json:"member"`
}
