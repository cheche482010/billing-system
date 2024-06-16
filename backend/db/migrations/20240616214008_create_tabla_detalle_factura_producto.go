package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTablaDetalleFacturaProducto, downCreateTablaDetalleFacturaProducto)
}

func upCreateTablaDetalleFacturaProducto(ctx context.Context, tx *sql.Tx) error {
	
	sqlStmt := `
	CREATE TABLE detalle_factura_producto (
		id INT AUTO_INCREMENT PRIMARY KEY,
		factura_id INT NOT NULL,
		producto_id INT NOT NULL,
		cantidad INT NOT NULL,
		precio_unitario DECIMAL(10, 2) NOT NULL,
		subtotal DECIMAL(10, 2) NOT NULL,
		FOREIGN KEY (factura_id) REFERENCES facturas(id),
		FOREIGN KEY (producto_id) REFERENCES productos(id),
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

func downCreateTablaDetalleFacturaProducto(ctx context.Context, tx *sql.Tx) error {
	
	sqlStmt := `
	DROP TABLE IF EXISTS detalle_factura_producto;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}
