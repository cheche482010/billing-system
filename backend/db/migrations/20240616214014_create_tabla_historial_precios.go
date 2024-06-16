package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTablaHistorialPrecios, downCreateTablaHistorialPrecios)
}

func upCreateTablaHistorialPrecios(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	CREATE TABLE historial_precios (
		id INT AUTO_INCREMENT PRIMARY KEY,
		producto_id INT NOT NULL,
		fecha DATE NOT NULL,
		precio DECIMAL(10, 2) NOT NULL,
		FOREIGN KEY (producto_id) REFERENCES productos(id),
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}


func downCreateTablaHistorialPrecios(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	DROP TABLE IF EXISTS historial_precios;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

