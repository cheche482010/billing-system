package controllers

import (
	"billing-system/models"
	"billing-system/services"
	"billing-system/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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
func (c *ClientesController) Create(w http.ResponseWriter, r *http.Request) {

	var cliente models.Cliente

	handleInvalidMethod(w, r, http.MethodPost)
	decodeRequestBody(w, r, &cliente)

	responseService := c.Service.Create(cliente)

	if responseService.Status {
		responseService.Message.Response["Message"] = "Cliente creado exitosamente."
		writeJSONResponse(w, responseService.Message, http.StatusCreated)
	} else {
		writeJSONResponse(w, responseService.Error, http.StatusBadRequest)
	}
}

// GetHandler handles GET requests to retrieve a cliente by ID
func (c *ClientesController) GetBy(w http.ResponseWriter, r *http.Request) {

	handleInvalidMethod(w, r, http.MethodGet)

	var cliente models.Cliente
	decodeRequestBody(w, r, &cliente)

	responseService := c.Service.Get(cliente.ID)

	if responseService.Status {
		successResponse := map[string]interface{}{
			"Code":   200,
			"Status": "Success",
			"Data":   responseService.Data,
		}
		writeJSONResponse(w, successResponse, http.StatusOK)
	} else {
		writeJSONResponse(w, responseService.Error, http.StatusInternalServerError)
	}
}

// GetAllHandler handles GET requests to retrieve all clientes
func (c *ClientesController) GetAll(w http.ResponseWriter, r *http.Request) {

	handleInvalidMethod(w, r, http.MethodGet)

	responseService := c.Service.GetAll()

	if responseService.Status {
		successResponse := map[string]interface{}{
			"Code":   200,
			"Status": "Success",
			"Data":   responseService.Data,
		}
		writeJSONResponse(w, successResponse, http.StatusOK)
	} else {
		writeJSONResponse(w, responseService.Error, http.StatusInternalServerError)
	}

}

// UpdateHandler handles PUT/PATCH requests to update a cliente
func (c *ClientesController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var cliente models.Cliente

	handleInvalidMethod(w, r, http.MethodPost)

	responseService := c.Service.Update(cliente)

	if responseService.Status {
		responseService.Message.Response["Message"] = "Cliente actualizado exitosamente."
		writeJSONResponse(w, responseService.Message, http.StatusCreated)
	} else {
		writeJSONResponse(w, responseService.Error, http.StatusBadRequest)
	}
}

// DeleteHandler handles DELETE requests to remove a cliente
func (c *ClientesController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	// if err := c.Service.Delete(id); err != nil {

}

func writeJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func handleInvalidMethod(w http.ResponseWriter, r *http.Request, allowedMethod string) {
	if r.Method != allowedMethod {
		responseData := utils.ErrorMethodMessage{
			Code:      http.StatusMethodNotAllowed,
			Status:    "Error",
			Message:   "Invalid request method",
			ErrorType: fmt.Sprintf("method %s not supported", r.Method),
			Timestamp: time.Now().Format(time.RFC3339),
		}
		json.NewEncoder(w).Encode(responseData)
		return
	}
}

func decodeRequestBody(w http.ResponseWriter, r *http.Request, target interface{}) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(target); err != nil {
		responseData := utils.ErrorJSONdMessage{
			Code:      http.StatusBadRequest,
			Status:    "Error",
			Message:   "Failed to decode JSON",
			ErrorType: fmt.Sprintf("Failed to decode JSON: %v", err),
			Timestamp: time.Now().Format(time.RFC3339),
		}
		json.NewEncoder(w).Encode(responseData)
		return
	}
}
