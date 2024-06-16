package clientes

import (
	"billing-system/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateModelCliente(t *testing.T) {

	cliente := models.Cliente{
		Nombre:    "Test Name",
		Apellido:  "Test Lastname",
		Cedula:    "123456789",
		Telefono:  "555-555-5555",
		Correo:    "test@example.com",
		Is_active: "true",
		CreatedAt: "2023-04-01T00:00:00Z",
		UpdatedAt: "2023-04-01T00:00:00Z",
	}

	err := Validate(cliente)

	assert.NoError(t, err)

}

func Validate(c models.Cliente) error {
	if c.Nombre == "" || c.Apellido == "" || c.Cedula == "" || c.Telefono == "" || c.Correo == "" {
		return errors.New("todos los campos son obligatorios")
	}
	return nil
}
