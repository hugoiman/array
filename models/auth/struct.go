package auth

type Session struct {
  Session      []DataSession  `json:"member"`
}
type DataSession struct {
  Id_member   int       `json:"id_member"`
  Nama        string    `json:"nama"`
  Slug        string    `json:"slug"`
}
