package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTablaProductos, downCreateTablaProductos)
}

func upCreateTablaProductos(ctx context.Context, tx *sql.Tx) error {

	sqlStmt := `
	CREATE TABLE productos (
		id INT AUTO_INCREMENT PRIMARY KEY,
		nombre VARCHAR(100) NOT NULL,
		precio DECIMAL(10, 2) NOT NULL,
		cantidad INT NOT NULL,
		codigo_barra VARCHAR(50) NOT NULL UNIQUE,
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

func downCreateTablaProductos(ctx context.Context, tx *sql.Tx) error {

	sqlStmt := `
	DROP TABLE IF EXISTS productos;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}
