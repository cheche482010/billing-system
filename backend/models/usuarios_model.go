package models

type Usuarios struct {
	ID         int    `json:"id"`
	Nombre     string `json:"nombre"`
	Contraseña string `json:"contraseña"`
	Tipo       string `json:"tipo"`
	IsActive   bool   `json:"is_active"`
	Blocked    bool   `json:"blocked"`
	Token      string `json:"token"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
