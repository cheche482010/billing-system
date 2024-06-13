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
func (r *ClientesRepository) Update(cliente models.Cliente) ResponseRepository {
	query := `
	UPDATE Clientes SET nombre = :nombre, apellido = :apellido, cedula = :cedula, telefono = :telefono, correo = :correo, updated_at = NOW()
	WHERE id = :id
	`
	result, err := db.DB.Exec(query, cliente)

	if err != nil {
		return ResponseRepository{Status: false, Error: ErrorMessage(err)}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return ResponseRepository{Status: false, Error: ErrorMessage(err)}
	}

	return ResponseRepository{Status: true, Data: rowsAffected}
}

// Delete removes a cliente by ID
func (r *ClientesRepository) Delete(id int) ResponseRepository {
	query := "DELETE FROM Clientes WHERE id =?"
	_, err := db.DB.Exec(query, id)
	if err != nil {
		return ResponseRepository{Status: false, Error: ErrorMessage(err)}
	}
	return ResponseRepository{Status: true}
}

func (r *ClientesRepository) SoftDelete(id int) error {

	client := db.StatusTable{
		TableName: "Clientes",
		ID:        1,
		IsActive:  false,
	}

	_, err := db.SwitchStatus(client)
	if err != nil {
		return err
	}

	return nil
}

func ErrorMessage(err error) utils.ErrorSqlMessage {
	return utils.ErrorSqlMessage{
		Code:      400,
		Message:   "Bad Request",
		ErroType:  err.Error(),
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
