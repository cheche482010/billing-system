package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTablaUsuarios, downCreateTablaUsuarios)
}

func upCreateTablaUsuarios(ctx context.Context, tx *sql.Tx) error {
	
	sqlStmt := `
	CREATE TABLE usuarios (
		id INT AUTO_INCREMENT PRIMARY KEY,
		nombre VARCHAR(100) NOT NULL,
		contraseña VARCHAR(255) NOT NULL,
		tipo ENUM('cajero', 'admin') NOT NULL,
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

func downCreateTablaUsuarios(ctx context.Context, tx *sql.Tx) error {
	
	sqlStmt := `
	DROP TABLE IF EXISTS usuarios;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}
