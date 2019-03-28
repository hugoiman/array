package structs

type Lokasi struct {
  Lokasi                []DataLokasi   `json:"lokasi"`
}
type DataLokasi struct {
  Id_lokasi             int            `json:"id_lokasi"`
  Cabang                string         `json:"cabang"`
  Alamat                string         `json:"alamat"`
}
