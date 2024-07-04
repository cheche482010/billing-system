package repositories

import (
	"billing-system/models"
	"database/sql"
	"errors"
)

type UsuariosRepository struct {
	DB *sql.DB
}

func NewUsuariosRepository(db *sql.DB) *UsuariosRepository {
	return &UsuariosRepository{DB: db}
}

// IniciarSesion verifica las credenciales del usuario y devuelve un token
func (r *UsuariosRepository) IniciarSesion(nombre, contraseña string) (models.Usuarios, error) {
	var usuario models.Usuarios
	query := "SELECT id, nombre, contraseña, tipo, is_active, token FROM usuarios WHERE nombre =? AND contraseña =?"
	err := r.DB.QueryRow(query, nombre, contraseña).Scan(&usuario.ID, &usuario.Nombre, &usuario.Contraseña, &usuario.Tipo, &usuario.IsActive, &usuario.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Usuarios{}, errors.New("nombre o contraseña incorrectos")
		}
		return models.Usuarios{}, err
	}

	return usuario, nil
}

type ResponseRepository struct {
	Status bool
	Data   interface{}
	Error  interface{}
}

package repositories

import (
	"billing-system/models"
	"database/sql"
	"errors"
	"time"

	// Asumiendo que tienes un paquete db donde se define Check
	"tu_paquete/db"
)

type UsuariosRepository struct {
	DB *sql.DB
}

func NewUsuariosRepository(db *sql.DB) *UsuariosRepository {
	return &UsuariosRepository{DB: db}
}

type ResponseRepository struct {
	Status bool
	Data   interface{}
	Error  interface{}
}

func (r *UsuariosRepository) IniciarSesion(nombre, contraseña string) ResponseRepository {
	var usuario models.Usuarios
	query := "SELECT id, nombre , tipo, is_active, token FROM usuarios WHERE nombre =? AND contraseña =?"
	err := r.DB.QueryRow(query, nombre, contraseña).Scan(&usuario.ID, &usuario.Nombre, &usuario.Contraseña, &usuario.Tipo, &usuario.IsActive, &usuario.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return ResponseRepository{Status: false, Data: nil, Error: "Nombre o contraseña incorrectos"}
		} else {
			return ResponseRepository{Status: false, Data: nil, Error: "Error interno del servidor"}
		}
	}

	isBlocked, err := db.Check(db.StatusTable{
		TableName: "usuarios",
		ID:        usuario.ID,
		Colum:     "blocked",
		Value:     false,
	})
	if err != nil || isBlocked {
		errorResponse := ErrorResponse{
			Code:       400,
			Message:    "Bad Request",
			ErroType:   "UsuarioBloqueado",
			Timestamp: time.Now().Format(time.RFC3339),
		}
		return ResponseRepository{Status: false, Data: nil, Error: errorResponse}
	}

	isActive, err := db.Check(db.StatusTable{
		TableName: "usuarios",
		ID:        usuario.ID,
		Colum:     "is_active",
		Value:     true,
	})
	if err != nil || !isActive {
		errorResponse := ErrorResponse{
			Code:       403,
			Message:    "Forbidden",
			ErroType:   "UsuarioDesactivado",
			Timestamp: time.Now().Format(time.RFC3339),
		}
		return ResponseRepository{Status: false, Data: nil, Error: errorResponse}
	}

	return ResponseRepository{Status: true, Data: usuario, Error: nil}
}

