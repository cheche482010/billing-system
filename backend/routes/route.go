package routes

import (
	"billing-system/controllers"
	"billing-system/db"
	"billing-system/repositories"
	"billing-system/services"
	"billing-system/utils"
	"encoding/json"
	"net/http"
	"regexp"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var Router = mux.NewRouter().PathPrefix("/api/").Subrouter()

func InitializeRoutes() error {
	database := db.GetDB()

	// Seguridad de rutas
	Router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	// Establecer modulos
	clienteRepo := repositories.NewClientesRepository(database)
	clienteService := services.NewClientesService(clienteRepo)
	clienteController := controllers.NewClientesController(clienteService)

	// Establecer no accesibles
	Router.HandleFunc("/clientes", handleForbidden).Methods("GET")

	// Establecer rutas permitidas
	Router.HandleFunc("/clientes/test", clienteController.Test).Methods("GET")
	Router.HandleFunc("/clientes/create", clienteController.Create).Methods("POST")

	// Manejadores de errores
	Router.NotFoundHandler = http.HandlerFunc(handleNotFound)
	Router.MethodNotAllowedHandler = http.HandlerFunc(handleMethodNotAllowed)
	Router.Use(validateURLFormat)

	return nil
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	response := utils.ErrorRouteMessage{
		Code:      http.StatusNotFound,
		Status:    "error",
		Error:     true,
		Message:   "Not Found",
		ErrorType: "The requested resource was not found on this server.",
		Path:      r.URL.Path,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	writeJSONResponse(w, response, http.StatusNotFound)
}

func handleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	response := utils.ErrorRouteMessage{
		Code:      http.StatusMethodNotAllowed,
		Status:    "error",
		Error:     true,
		Message:   "Method Not Allowed",
		ErrorType: "The requested method is not allowed for the requested resource.",
		Path:      r.URL.Path,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	writeJSONResponse(w, response, http.StatusMethodNotAllowed)
}

func handleBadRequest(w http.ResponseWriter, r *http.Request) {
	response := utils.ErrorRouteMessage{
		Code:      http.StatusBadRequest,
		Status:    "error",
		Error:     true,
		Message:   "Bad Request",
		ErrorType: "The requested URL format is not valid.",
		Path:      r.URL.Path,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	writeJSONResponse(w, response, http.StatusBadRequest)
}

func handleForbidden(w http.ResponseWriter, r *http.Request) {
	response := utils.ErrorRouteMessage{
		Code:      http.StatusForbidden,
		Status:    "error",
		Error:     true,
		Message:   "Forbidden",
		ErrorType: "You do not have the necessary permissions to access the resource.",
		Path:      r.URL.Path,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	writeJSONResponse(w, response, http.StatusForbidden)
}

func writeJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func checkURLFormat(url string) bool {
	pattern := "^/api/[A-Za-z]+(/[A-Za-z]+)*$"
	match, _ := regexp.MatchString(pattern, url)
	return match
}

func validateURLFormat(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !checkURLFormat(r.URL.Path) {
			handleBadRequest(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
