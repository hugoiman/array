package structs

import(
  "time"
)

type SessionMember struct {
  Id_member   int       `json:"id_member"`
  Nama        string    `json:"nama"`
  Slug        string    `json:"slug"`
}

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

  CustTgl_lahir         string
  CustTgl_gabung        string
  Lokasi                DataLokasi     `json:"lokasi"`
}

type DataLokasi struct {
  Id_lokasi             int            `json:"id_lokasi"`
  Cabang                string         `json:"cabang"`
  Alamat                string         `json:"alamat"`
}
