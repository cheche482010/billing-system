package models

type Cliente struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Cedula    string `json:"cedula"`
	Telefono  string `json:"telefono"`
	Correo    string `json:"correo"`
	Is_active string `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
