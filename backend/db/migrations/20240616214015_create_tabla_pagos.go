package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTablaPagos, downCreateTablaPagos)
}

func upCreateTablaPagos(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	CREATE TABLE pagos (
		id INT AUTO_INCREMENT PRIMARY KEY,
		cliente_factura_id INT NOT NULL,
		fecha_pago DATE NOT NULL,
		pago_total DECIMAL(10, 2) NOT NULL,
		FOREIGN KEY (cliente_factura_id) REFERENCES cliente_facturas(id),
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}


func downCreateTablaPagos(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	DROP TABLE IF EXISTS pagos;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

