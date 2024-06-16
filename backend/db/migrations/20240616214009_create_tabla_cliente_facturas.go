package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTablaClienteFacturas, downCreateTablaClienteFacturas)
}

func upCreateTablaClienteFacturas(ctx context.Context, tx *sql.Tx) error {
	
	sqlStmt := `
	CREATE TABLE cliente_facturas (
		id INT AUTO_INCREMENT PRIMARY KEY,
		cliente_id INT NOT NULL,
		factura_id INT NOT NULL,
		fecha_compra DATE NOT NULL,
		FOREIGN KEY (cliente_id) REFERENCES clientes(id),
		FOREIGN KEY (factura_id) REFERENCES facturas(id),
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

func downCreateTablaClienteFacturas(ctx context.Context, tx *sql.Tx) error {
	
	sqlStmt := `
	DROP TABLE IF EXISTS cliente_facturas;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}
