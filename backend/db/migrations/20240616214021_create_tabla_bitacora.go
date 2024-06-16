package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTablaBitacora, downCreateTablaBitacora)
}

func upCreateTablaBitacora(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	CREATE TABLE bitacora (
		id INT AUTO_INCREMENT PRIMARY KEY,
		accion VARCHAR(255) NOT NULL,
		descripcion TEXT,
		tipo ENUM('INFO', 'ERROR', 'WARNING') NOT NULL,
		fecha_hora TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		usuario_id INT,
		FOREIGN KEY (usuario_id) REFERENCES usuarios(id),
		is_active BOOLEAN DEFAULT TRUE
	);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}


func downCreateTablaBitacora(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	DROP TABLE IF EXISTS bitacora;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

