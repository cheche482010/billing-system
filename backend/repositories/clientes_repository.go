package repositories

import (
	"billing-system/db"
	"billing-system/models"
	"billing-system/utils"
	"database/sql"
	"time"
)

type ClientesRepository struct {
	DB *sql.DB
}

func NewClientesRepository(db *sql.DB) *ClientesRepository {
	return &ClientesRepository{DB: db}
}

type ResponseRepository struct {
	Status bool
	ID     int64
	Error  utils.ErrorSqlMessage
}

// Create inserts a new cliente into the database
func (r *ClientesRepository) Create(cliente models.Cliente) ResponseRepository {
	query := `
	INSERT INTO Clientes (nombre, apellido, cedula, telefono, correo, created_at, updated_at)
	VALUES (?,?,?,?,?, NOW(), NOW())
	`
	result, err := db.DB.Exec(query, cliente.Nombre, cliente.Apellido, cliente.Cedula, cliente.Telefono, cliente.Correo)

	if err != nil {
		return ResponseRepository{Status: false, Error: ErrorMessage(err)}
	} else {
		lastID, err := result.LastInsertId()
		if err != nil {
			return ResponseRepository{Status: false, Error: ErrorMessage(err)}
		}
		return ResponseRepository{Status: true, ID: lastID}
	}
}

// Get retrieves a single cliente by ID
func (r *ClientesRepository) Get(id int) (*models.Cliente, error) {
	cliente := &models.Cliente{}
	query := "SELECT id, nombre, apellido, cedula, telefono, correo, created_at, updated_at FROM Clientes WHERE id =?"
	err := db.DB.QueryRow(query, id).Scan(&cliente.ID, &cliente.Nombre, &cliente.Apellido, &cliente.Cedula, &cliente.Telefono, &cliente.Correo, &cliente.CreatedAt, &cliente.UpdatedAt)
	return cliente, err
}

// GetAll retrieves all clientes
func (r *ClientesRepository) GetAll() ([]*models.Cliente, error) {
	clientes := []*models.Cliente{}
	query := "SELECT id, nombre, apellido, cedula, telefono, correo, created_at, updated_at FROM Clientes"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cliente models.Cliente
		if err := rows.Scan(&cliente.ID, &cliente.Nombre, &cliente.Apellido, &cliente.Cedula, &cliente.Telefono, &cliente.Correo, &cliente.CreatedAt, &cliente.UpdatedAt); err != nil {
			return nil, err
		}
		clientes = append(clientes, &cliente)
	}
	return clientes, nil
}

// Update updates a cliente's details
func (r *ClientesRepository) Update(cliente models.Cliente) error {
	query := `
	UPDATE Clientes SET nombre =?, apellido =?, cedula =?, telefono =?, correo =?, updated_at = NOW()
	WHERE id =?
	`
	_, err := db.DB.Exec(query, cliente.Nombre, cliente.Apellido, cliente.Cedula, cliente.Telefono, cliente.Correo, cliente.ID)
	return err
}

// Delete removes a cliente by ID
func (r *ClientesRepository) Delete(id int) error {
	query := "DELETE FROM Clientes WHERE id =?"
	_, err := db.DB.Exec(query, id)
	return err
}

func ErrorMessage(err error) utils.ErrorSqlMessage {
	return utils.ErrorSqlMessage{
		Code:      400,
		Message:   "Bad Request",
		ErroType:  err.Error(),
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
