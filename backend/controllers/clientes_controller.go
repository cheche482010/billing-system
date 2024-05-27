package controllers

import (
	"billing-system/models"
	"billing-system/services"
	"encoding/json"
	"net/http"
)

type ClientesController struct {
	Service *services.ClientesService
}

func NewClientesController(service *services.ClientesService) *ClientesController {
	return &ClientesController{Service: service}
}

func (c *ClientesController) Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola Mundo"))
}

// CreateHandler handles POST requests to create a new cliente
func (c *ClientesController) CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var cliente models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	// if err := c.Service.Create(cliente); err != nil {

	// 	var errorValidationMessage utils.ErrorValidationMessage
	// 	if err = json.Unmarshal([]byte(err.Error()), &errorValidationMessage); err == nil {
	// 		w.Header().Set("Content-Type", "application/json")
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		json.NewEncoder(w).Encode(errorValidationMessage)
	// 	} else {
	// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
	// 	}
	// 	return
	// }

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cliente)
}

// GetHandler handles GET requests to retrieve a cliente by ID
func (c *ClientesController) GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// id := chi.URLParam(r, "id")

	// cliente, err := c.Service.Get(id)
	// if err != nil {
	// 	http.Error(w, "Failed to get cliente", http.StatusInternalServerError)
	// 	return
	// }

	// json.NewEncoder(w).Encode(cliente)
}

// GetAllHandler handles GET requests to retrieve all clientes
func (c *ClientesController) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	clientes, err := c.Service.GetAll()
	if err != nil {
		http.Error(w, "Failed to get all clientes", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(clientes)
}

// UpdateHandler handles PUT/PATCH requests to update a cliente
func (c *ClientesController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var cliente models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	if err := c.Service.Update(cliente); err != nil {
		http.Error(w, "Failed to update cliente", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cliente)
}

// DeleteHandler handles DELETE requests to remove a cliente
func (c *ClientesController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	// id := chi.URLParam(r, "id")

	// if err := c.Service.Delete(id); err != nil {
	// 	http.Error(w, "Failed to delete cliente", http.StatusInternalServerError)
	// 	return
	// }

	// w.WriteHeader(http.StatusNoContent)
}
