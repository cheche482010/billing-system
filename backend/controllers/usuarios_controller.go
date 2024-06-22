package controllers

import (
	"billing-system/services"
	"encoding/json"
	"net/http"
)

type AutenticacionController struct {
	Service *services.AutenticacionService
}

func NewAutenticacionController(service *services.AutenticacionService) *AutenticacionController {
	return &AutenticacionController{Service: service}
}

func (c *AutenticacionController) Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Nombre     string `json:"nombre"`
		Contraseña string `json:"contraseña"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	usuario, err := c.Service.IniciarSesion(creds.Nombre, creds.Contraseña)
	if err != nil {
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuario)
}
