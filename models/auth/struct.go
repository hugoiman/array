package auth

type SessionMember struct {
  Id_member   int       `json:"id_member"`
  Nama        string    `json:"nama"`
  Slug        string    `json:"slug"`
}

type SessionAdmin struct {
  Id_admin    int       `json:"id_admin"`
  Nama        string    `json:"nama"`
  Slug        string    `json:"slug"`
}
