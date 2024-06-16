package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTablaExample, downCreateTablaExample)
}

// upCreateTablaExample crea la tabla "example".
func upCreateTablaExample(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	CREATE TABLE example (
		id INT AUTO_INCREMENT PRIMARY KEY,
		example VARCHAR(100) NOT NULL,
	);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

// downCreateTablaExample elimina la tabla "example".
func downCreateTablaExample(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	DROP TABLE IF EXISTS example;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

// upAddColumnToExample agrega una nueva columna a la tabla "example".
func upAddColumnToExample(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	ALTER TABLE example ADD COLUMN description TEXT;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

// downAddColumnToExample elimina la columna recién agregada de la tabla "example".
func downAddColumnToExample(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	ALTER TABLE example DROP COLUMN description;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

// upChangeColumnTypeInExample cambia el tipo de una columna en la tabla "example".
func upChangeColumnTypeInExample(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	ALTER TABLE example MODIFY example VARCHAR(200);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

// downChangeColumnTypeInExample reestablece el tipo de columna original en la tabla "example".
func downChangeColumnTypeInExample(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	ALTER TABLE example MODIFY example VARCHAR(100);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

// upRemoveForeignKeyFromExample elimina una restricción de clave foránea de la tabla "example".
func upRemoveForeignKeyFromExample(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	ALTER TABLE example DROP FOREIGN KEY fk_example_description;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

// downRemoveForeignKeyFromExample agrega la restricción de clave foránea eliminada a la tabla "example".
func downRemoveForeignKeyFromExample(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	ALTER TABLE example ADD CONSTRAINT fk_example_description FOREIGN KEY (description) REFERENCES descriptions(id);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

// upCreateIndexOnExample crea un índice en la tabla "example".
func upCreateIndexOnExample(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	CREATE INDEX idx_example_example ON example(example);
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

// downCreateIndexOnExample elimina el índice creado en la tabla "example".
func downCreateIndexOnExample(ctx context.Context, tx *sql.Tx) error {
	sqlStmt := `
	DROP INDEX idx_example_example;
	`
	_, err := tx.ExecContext(ctx, sqlStmt)
	return err
}

func upInsertDatosEnEjemploDatos(ctx context.Context, tx *sql.Tx) error {
	// Definición de los datos a insertar.
	datos := []struct {
		nombre    string
		valor     string
	}{
		{"Dato 1", "Valor 1"},
		{"Dato 2", "Valor 2"},
		{"Dato 3", "Valor 3"},
	}

	// Preparación del statement para insertar datos.
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO example (nombre, valor) VALUES (?,?)")
	if err!= nil {
		return err
	}
	defer stmt.Close()

	// Iterar sobre los datos y realizar las inserciones.
	for _, dato := range datos {
		_, err := stmt.ExecContext(ctx, dato.nombre, dato.valor)
		if err!= nil {
			return err
		}
	}

	return nil
}
