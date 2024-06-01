package services

import (
	"billing-system/models"
	"billing-system/repositories"
	"billing-system/utils"
	"time"
)

type ClientesService struct {
	repo *repositories.ClientesRepository
}

func NewClientesService(repo *repositories.ClientesRepository) *ClientesService {
	return &ClientesService{repo: repo}
}

type ResponseService struct {
	Status  bool
	Error   interface{}
	Data    interface{}
	Message Response
}

type Response struct {
	Code      int                    `json:"code"`
	Status    string                 `json:"status"`
	Message   string                 `json:"message"`
	Error     bool                   `json:"error"`
	Timestamp string                 `json:"timestamp"`
	Response  map[string]interface{} `json:"response"`
}

func (s *ClientesService) Create(cliente models.Cliente) ResponseService {

	validationRules := map[string]utils.ValidationCriteria{
		"nombre": {
			Require:   true,
			Pattern:   utils.Regex["caracteres"],
			MinLength: 3,
			MaxLength: 20,
		},
	}

	validation := utils.ValidateData(cliente.Nombre, validationRules)

	if validation.Status {
		response := s.repo.Create(cliente)

		if response.Status {
			resp := Response{
				Code:      200,
				Status:    "success",
				Message:   "Request successfully processed",
				Error:     false,
				Timestamp: time.Now().Format(time.RFC3339),
				Response: map[string]interface{}{
					"Code": 201,
					"Data": map[string]interface{}{
						"id": response.Data,
					},
					"Message": "",
				},
			}
			return ResponseService{Status: true, Message: resp}
		} else {
			return ResponseService{Status: false, Error: response.Error}
		}
	} else {
		return ResponseService{Status: false, Error: validation.Error}
	}

}

// Get wraps the repository's Get method
func (s *ClientesService) Get(id int) ResponseService {

	validationRules := map[string]utils.ValidationCriteria{
		"id": {
			Require:   true,
			IsInteger: true,
		},
	}

	validation := utils.ValidateData(id, validationRules)

	if validation.Status {
		response := s.repo.Get(id)
		return ResponseService{Status: true, Data: response.Data}
	} else {
		return ResponseService{Status: false, Error: validation.Error}
	}
}

// GetAll wraps the repository's GetAll method
func (s *ClientesService) GetAll() repositories.ResponseRepository {
	return s.repo.GetAll()
}

// Update wraps the repository's Update method
func (s *ClientesService) Update(cliente models.Cliente) ResponseService {

	validationRules := map[string]utils.ValidationCriteria{
		"nombre": {
			Require:   true,
			Pattern:   utils.Regex["caracteres"],
			MinLength: 3,
			MaxLength: 20,
		},
	}

	validation := utils.ValidateData(cliente.Nombre, validationRules)

	if validation.Status {
		response := s.repo.Update(cliente)

		if response.Status {
			resp := Response{
				Code:      200,
				Status:    "success",
				Message:   "Request successfully processed",
				Error:     false,
				Timestamp: time.Now().Format(time.RFC3339),
				Response: map[string]interface{}{
					"Code": 201,
					"Data": map[string]interface{}{
						"id": response.Data,
					},
					"Message": "",
				},
			}
			return ResponseService{Status: true, Message: resp}
		} else {
			return ResponseService{Status: false, Error: response.Error}
		}
	} else {
		return ResponseService{Status: false, Error: validation.Error}
	}
}

// Delete wraps the repository's Delete method
func (s *ClientesService) Delete(id int) error {
	return s.repo.Delete(id)
}
