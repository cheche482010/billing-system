package services

import (
	"billing-system/models"
	"billing-system/repositories"
)

type AutenticacionService struct {
	repo *repositories.UsuariosRepository
}

func NewAutenticacionService(repo *repositories.UsuariosRepository) *AutenticacionService {
	return &AutenticacionService{repo: repo}
}

// IniciarSesion utiliza el repositorio para iniciar sesión
func (s *AutenticacionService) IniciarSesion(nombre, contraseña string) (models.Usuarios, error) {
	return s.repo.IniciarSesion(nombre, contraseña)
}
