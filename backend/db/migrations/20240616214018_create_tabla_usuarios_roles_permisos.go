package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTablaUsuariosRolesPermisos, downCreateTablaUsuariosRolesPermisos)
}

func upCreateTablaUsuariosRolesPermisos(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	CREATE TABLE usuarios_roles_permisos (
		id INT AUTO_INCREMENT PRIMARY KEY,
		usuario_id INT NOT NULL,
		rol VARCHAR(20) NOT NULL,
		permiso VARCHAR(255) NOT NULL,
		FOREIGN KEY (usuario_id) REFERENCES usuarios(id),
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}


func downCreateTablaUsuariosRolesPermisos(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	DROP TABLE IF EXISTS usuarios_roles_permisos;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

