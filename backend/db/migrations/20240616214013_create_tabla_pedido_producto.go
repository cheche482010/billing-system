package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTablaPedidoProducto, downCreateTablaPedidoProducto)
}

func upCreateTablaPedidoProducto(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	CREATE TABLE pedido_producto (
		id INT AUTO_INCREMENT PRIMARY KEY,
		pedido_id INT NOT NULL,
		producto_id INT NOT NULL,
		cantidad INT NOT NULL,
		FOREIGN KEY (pedido_id) REFERENCES pedidos(id),
		FOREIGN KEY (producto_id) REFERENCES productos(id),
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}


func downCreateTablaPedidoProducto(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	DROP TABLE IF EXISTS pedido_producto;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

