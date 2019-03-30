package structs

import(
  "time"
)

type SessionAdmin struct {
  Id_admin    int       `json:"id_admin"`
  Nama        string    `json:"nama"`
  Slug        string    `json:"slug"`
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

  CustTgl_lahir         string          `json:custTgl_lahir`
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
