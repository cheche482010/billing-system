package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTablaMetodoPago, downCreateTablaMetodoPago)
}

func upCreateTablaMetodoPago(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	CREATE TABLE metodo_pago (
		id INT AUTO_INCREMENT PRIMARY KEY,
		pago_id INT NOT NULL,
		metodo_pago VARCHAR(50) NOT NULL,
		monto DECIMAL(10, 2) NOT NULL,
		FOREIGN KEY (pago_id) REFERENCES pagos(id),
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}


func downCreateTablaMetodoPago(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	DROP TABLE IF EXISTS metodo_pago;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

