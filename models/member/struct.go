package member

type Member struct {
  Member           []DataMember   `json:"member"`
}
type DataMember struct {
  Id_member        int            `json:"id_member"`
  Id_lokasi        int            `json:"id_lokasi"`
  Nama             string         `json:"nama"`
  Email            string         `json:"email"`
  Password         string         `json:"password"`
  Nik              string         `json:"nik"`
  No_hp            string         `json:"no_hp"`
  Foto             string         `json:"foto"`
  Tgl_lahir        string         `json:"tgl_lahir"`
  Pekerjaan        string         `json:"pekerjaan"`
  Alamat_asal      string         `json:"alamat_asal"`
  No_kamar         string         `json:"no_kamar"`
  Tgl_gabung       string         `json:"tgl_gabung"`
  Status_member    string         `json:"status_member"`
  Tipe_pembayaran  string         `json:"tipe_pembayaran"`
  Username_wifi    string         `json:"username_wifi"`
  Slug             string         `json:"slug"`
  Universitas      string         `json:"universitas"`
  Jurusan          string         `json:"jurusan"`
  Angkatan         string         `json:"angkatan"`
  Nama_kerabat1    string         `json:"nama_kerabat1"`
  Nama_kerabat2    string         `json:"nama_kerabat2"`
  Hubungan1        string         `json:"hubungan1"`
  Hubungan2        string         `json:"hubungan2"`
  No_hp1           string         `json:"no_hp1"`
  No_hp2           string         `json:"no_hp2"`
  No_stnk          string         `json:"no_stnk"`
  Perusahaan       string         `json:"Perusahaan"`
  Jabatan          string         `json:"jabatan"`
}

type Informasi struct {
  Informasi        []DataInformasi     `json:"informasi"`
}

type DataInformasi struct {
  Id_informasi     int            `json:"id_informasi"`
  Judul            string         `json:"judul"`
  Keterangan       string         `json:"keterangan"`
  Foto             string         `json:"foto"`
  Waktu            string         `json:"waktu"`
  Status           string         `json:"status"`
}

type Lokasi struct {
  Lokasi           []DataLokasi   `json:"lokasi"`
}
type DataLokasi struct {
  Id_lokasi        int            `json:"id_lokasi"`
  Cabang           string         `json:"cabang"`
  Alamat           string         `json:"alamat"`
}
