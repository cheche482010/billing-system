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
	query := "SELECT id, nombre, contraseña, tipo, is_active, token, created_at, updated_at FROM usuarios WHERE nombre =? AND contraseña =?"
	err := r.DB.QueryRow(query, nombre, contraseña).Scan(&usuario.ID, &usuario.Nombre, &usuario.Contraseña, &usuario.Tipo, &usuario.IsActive, &usuario.Token, &usuario.CreatedAt, &usuario.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Usuarios{}, errors.New("nombre o contraseña incorrectos")
		}
		return models.Usuarios{}, err
	}

	return usuario, nil
}
