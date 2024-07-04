package models

type Usuarios struct {
	ID         int    `json:"id"`
	Nombre     string `json:"nombre"`
	Contraseña string `json:"contraseña"`
	Tipo       string `json:"tipo"`
	IsActive   bool   `json:"is_active"`
	Blocked    bool   `json:"blocked"`
	Token      string `json:"token"`
	SessionToken  string    `json:"session_token"`
	IsLogged  bool       `json:"is_logged"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
