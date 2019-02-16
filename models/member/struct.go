package member

type Member struct {
  Member      []DataMember  `json:"member"`
}
type DataMember struct {
  Id_member   int       `json:"id_member"`
  Nama        string    `json:"nama"`
  Email       string    `json:"email"`
  Password    string    `json:"password"`
}

type Informasi struct {
  Informasi []DataInformasi   `json:"informasi"`
}

type DataInformasi struct {
  Id_informasi    int         `json:"id_informasi"`
  Judul           string      `json:"judul"`
  Keterangan      string      `json:"keterangan"`
  Foto            string      `json:"foto"`
  Waktu           string      `json:"waktu"`
  Status          string      `json:"status"`
}

type Lokasi struct {
  Lokasi []DataLokasi   `json:"lokasi"`
}
type DataLokasi struct {
  Id_lokasi    int         `json:"id_lokasi"`
  Cabang       string      `json:"cabang"`
  Alamat       string      `json:"alamat"`
}
