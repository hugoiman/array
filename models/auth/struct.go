package auth

type Member struct {
  Member      []DataMember  `json:"member"`
}
type DataMember struct {
  Id_member   int       `json:"id_member"`
  Nama        string    `json:"nama"`
  Email       string    `json:"email"`
  Password    string    `json:"password"`
}
