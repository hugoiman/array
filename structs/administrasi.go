package structs

import(
  "time"
)

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

  // Member                DataMember  `json:"member"`
  CustCheck_in          string
  CustCheck_out         string
  CustTgl_pembayaran    string
}

type InfoTagihan struct {
  InfoTagihan           []DataTagihan `json:"info_tagihan"`
}
type DataTagihan struct {
  Check_in              time.Time   `json:"check_in"`
  Check_out             time.Time   `json:"check_out"`
  Tagihan               int         `json:"tagihan"`

  CustCheck_in          string
  CustCheck_out         string
}
