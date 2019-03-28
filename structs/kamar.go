package structs

type Kamar struct {
  Kamar                []DataKamar     `json:"kamar"`
}
type DataKamar struct {
  Id_kamar              int            `json:"id_kamar"`
  Id_lokasi             int            `json:"id_lokasi"`
  No_kamar              string         `json:"no_kamar"`
  Status_kamar          string         `json:"status_kamar"`
}
