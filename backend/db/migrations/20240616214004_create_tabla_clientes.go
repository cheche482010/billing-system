package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTablaClientes, downCreateTablaClientes)
}

func upCreateTablaClientes(ctx context.Context, tx *sql.Tx) error {
	
	sqlStmt := `
	CREATE TABLE clientes (
		id INT AUTO_INCREMENT PRIMARY KEY,
		nombre VARCHAR(100) NOT NULL,
		apellido VARCHAR(100) NOT NULL,
		cedula VARCHAR(20) NOT NULL UNIQUE,
		telefono VARCHAR(15),
		correo VARCHAR(100),
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

func downCreateTablaClientes(ctx context.Context, tx *sql.Tx) error {
	
	sqlStmt := `
	DROP TABLE IF EXISTS clientes;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}
