package services

import (
	"billing-system/models"
	"billing-system/repositories"
	"math/rand"
	"time"
	"crypto/rand"
	"encoding/base64"
	"crypto/sha256"
)

type AutenticacionService struct {
	repo *repositories.UsuariosRepository
}

func NewAutenticacionService(repo *repositories.UsuariosRepository) *AutenticacionService {
	return &AutenticacionService{repo: repo}
}

func generateSessionToken() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}

	hashedBytes := sha256.Sum256(bytes)
	token := base64.RawURLEncoding.EncodeToString(hashedBytes[:])

	return token
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
