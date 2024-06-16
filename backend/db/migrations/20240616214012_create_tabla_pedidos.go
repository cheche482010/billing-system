package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTablaPedidos, downCreateTablaPedidos)
}

func upCreateTablaPedidos(ctx context.Context, tx *sql.Tx) error {
	
	sqlStmt := `
	CREATE TABLE pedidos (
		id INT AUTO_INCREMENT PRIMARY KEY,
		cliente_id INT NOT NULL,
		fecha DATE NOT NULL,
		total DECIMAL(10, 2) NOT NULL,
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (cliente_id) REFERENCES clientes(id)
	);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

func downCreateTablaPedidos(ctx context.Context, tx *sql.Tx) error {
	
	sqlStmt := `
	DROP TABLE IF EXISTS pedidos;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}
