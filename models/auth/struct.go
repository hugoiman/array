package auth

type Member struct {
  Member      []DataMember  `json:"member"`
}
type DataMember struct {
  Id_member   int       `json:"id_member"`
  Nama        string    `json:"nama"`
  Slug        string    `json:"slug"`
  Password    string    `json:"password"`
}
