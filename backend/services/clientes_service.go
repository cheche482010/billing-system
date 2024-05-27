package services

import (
	"billing-system/models"
	"billing-system/repositories"
	"billing-system/utils"
	"fmt"
	"time"
)

type ClientesService struct {
	repo *repositories.ClientesRepository
}

func NewClientesService(repo *repositories.ClientesRepository) *ClientesService {
	return &ClientesService{repo: repo}
}

type ResponseService struct {
	Success bool
	Error   utils.ErrorSqlMessage
	Message ResponseSuccess
}

type ResponseSuccess struct {
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

	result, err := utils.ValidateData(cliente.Nombre, validationRules)

	if err != nil {
		return fmt.Errorf("%w: %v", err, result)
	}

	response := s.repo.Create(cliente)

	if response.Status {
		resp := ResponseSuccess{
			Code:      200,
			Status:    "success",
			Message:   "Request successfully processed",
			Error:     false,
			Timestamp: time.Now().Format(time.RFC3339),
			Response: map[string]interface{}{
				"Code": 201,
				"Data": map[string]interface{}{
					"id": response.ID,
				},
				"Message": "",
			},
		}
		return ResponseService{Success: true, Message: resp}
	} else {
		return ResponseService{Success: false, Error: response.Error}
	}

}

// Get wraps the repository's Get method
func (s *ClientesService) Get(id int) (*models.Cliente, error) {
	return s.repo.Get(id)
}

// GetAll wraps the repository's GetAll method
func (s *ClientesService) GetAll() ([]*models.Cliente, error) {
	return s.repo.GetAll()
}

// Update wraps the repository's Update method
func (s *ClientesService) Update(cliente models.Cliente) error {
	return s.repo.Update(cliente)
}

// Delete wraps the repository's Delete method
func (s *ClientesService) Delete(id int) error {
	return s.repo.Delete(id)
}
