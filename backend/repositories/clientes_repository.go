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
	Data   interface{}
	Error  interface{}
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
		return ResponseRepository{Status: true, Data: lastID}
	}
}

// Get retrieves a single cliente by ID
func (r *ClientesRepository) Get(id int) ResponseRepository {
	cliente := &models.Cliente{}
	query := "SELECT id, nombre, apellido, cedula, telefono, correo, created_at, updated_at FROM Clientes WHERE id =?"
	err := db.DB.QueryRow(query, id).Scan(&cliente.ID, &cliente.Nombre, &cliente.Apellido, &cliente.Cedula, &cliente.Telefono, &cliente.Correo, &cliente.CreatedAt, &cliente.UpdatedAt)
	if err != nil {
		return ResponseRepository{Status: false, Error: ErrorMessage(err)}
	}
	return ResponseRepository{Status: true, Data: cliente, Error: false}
}

// GetAll retrieves all clientes
func (r *ClientesRepository) GetAll() ResponseRepository {
	clientes := []*models.Cliente{}
	query := "SELECT id, nombre, apellido, cedula, telefono, correo, created_at, updated_at FROM Clientes ORDER BY apellido ASC"
	rows, err := db.DB.Query(query)

	if err != nil {
		return ResponseRepository{Status: false, Error: ErrorMessage(err)}
	}

	defer rows.Close()

	for rows.Next() {
		var cliente models.Cliente
		if err := rows.Scan(&cliente.ID, &cliente.Nombre, &cliente.Apellido, &cliente.Cedula, &cliente.Telefono, &cliente.Correo, &cliente.CreatedAt, &cliente.UpdatedAt); err != nil {
			return ResponseRepository{Status: false, Error: ErrorMessage(err)}
		}
		clientes = append(clientes, &cliente)
	}
	if err := rows.Err(); err != nil {
		return ResponseRepository{Status: false, Error: ErrorMessage(err)}
	}

	data := interface{}(clientes)

	return ResponseRepository{Status: true, Data: data, Error: false}
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
