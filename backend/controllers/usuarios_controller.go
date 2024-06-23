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

	authResponse, err := c.Service.IniciarSesion(creds.Nombre, creds.Contraseña)
	response := map[string]interface{}{
		"status": authResponse.Status,
		"data":   authResponse.Data,
	}

	httpStatus := http.StatusOK
	if err != nil {
		response["message"] = "Login failed: " + err.Error()
		httpStatus = http.StatusUnauthorized
	} else {
		response["message"] = "Login successful"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(response)
}
