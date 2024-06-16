package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTablaProveedores, downCreateTablaProveedores)
}

func upCreateTablaProveedores(ctx context.Context, tx *sql.Tx) error {
	
	sqlStmt := `
	CREATE TABLE proveedores (
		id INT AUTO_INCREMENT PRIMARY KEY,
		nombre VARCHAR(100) NOT NULL,
		contacto VARCHAR(100),
		direccion VARCHAR(255),
		telefono VARCHAR(15),
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

func downCreateTablaProveedores(ctx context.Context, tx *sql.Tx) error {
	
	sqlStmt := `
	DROP TABLE IF EXISTS proveedores;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}
