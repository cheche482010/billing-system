package controllers

import (
	"billing-system/db"
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Code      int    `json:"code"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	Error     bool   `json:"error"`
	Timestamp string `json:"timestamp"`
	ErroType  string `json:"error_type,omitempty"`
}

func MigrateDB(w http.ResponseWriter, r *http.Request) {
	err := db.RunMigrations()

	var response Response
	var status int

	if err != nil {
		status = http.StatusInternalServerError
		response = Response{
			Code:      status,
			Status:    "error",
			Message:   "Error al ejecutar migraciones",
			Error:     true,
			Timestamp: time.Now().Format(time.RFC3339),
			ErroType:  err.Error(),
		}
	} else {
		status = http.StatusOK
		response = Response{
			Code:      status,
			Status:    "success",
			Message:   "Migraciones ejecutadas con éxito",
			Error:     false,
			Timestamp: time.Now().Format(time.RFC3339),
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
