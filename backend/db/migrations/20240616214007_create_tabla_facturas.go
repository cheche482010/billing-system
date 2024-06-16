package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTablaFacturas, downCreateTablaFacturas)
}

func upCreateTablaFacturas(ctx context.Context, tx *sql.Tx) error {
	
	sqlStmt := `
	CREATE TABLE facturas (
		id INT AUTO_INCREMENT PRIMARY KEY,
		num_factura VARCHAR(6) NOT NULL UNIQUE,
		fecha DATE NOT NULL,
		total DECIMAL(10, 2) NOT NULL,
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

func downCreateTablaFacturas(ctx context.Context, tx *sql.Tx) error {
	
	sqlStmt := `
	DROP TABLE IF EXISTS facturas;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}
