package services

import (
	"billing-system/models"
	"billing-system/repositories"
	"math/rand"
	"time"
)

type AutenticacionService struct {
	repo *repositories.UsuariosRepository
}

func NewAutenticacionService(repo *repositories.UsuariosRepository) *AutenticacionService {
	return &AutenticacionService{repo: repo}
}

func generateSessionToken() string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type AutenticacionResponse struct {
	Status bool              `json:"status"`
	Data   map[string]string `json:"data"`
}

// IniciarSesion utiliza el repositorio para iniciar sesión
func (s *AutenticacionService) IniciarSesion(nombre, contraseña string) (AutenticacionResponse, error) {
	usuario, err := s.repo.IniciarSesion(nombre, contraseña)
	if err != nil {
		return AutenticacionResponse{
			Status: false,
			Data:   nil,
		}, err
	}
	sesionToken := generateSessionToken()
	data := map[string]string{
		"user_token":   usuario.Token,
		"sesion_token": sesionToken,
	}
	return AutenticacionResponse{
		Status: true,
		Data:   data,
	}, nil
}
