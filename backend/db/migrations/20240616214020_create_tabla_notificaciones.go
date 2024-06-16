package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTablaNotificaciones, downCreateTablaNotificaciones)
}

func upCreateTablaNotificaciones(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	CREATE TABLE notificaciones (
		id INT AUTO_INCREMENT PRIMARY KEY,
		destinatario_id INT NOT NULL,
		mensaje TEXT NOT NULL,
		fecha_notificacion DATE NOT NULL,
		leido BOOLEAN DEFAULT FALSE,
		FOREIGN KEY (destinatario_id) REFERENCES usuarios(id),
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}


func downCreateTablaNotificaciones(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	DROP TABLE IF EXISTS notificaciones;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

