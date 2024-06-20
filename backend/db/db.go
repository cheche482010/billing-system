package db

import (
	"database/sql"
	"fmt"
	"time"

	"billing-system/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	jsoniter "github.com/json-iterator/go"
)

var DB *sql.DB

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Connect() bool {
	var err error
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		User, Passwd, Host, Port, Database)
	DB, err = sql.Open("mysql", connection)
	if err != nil {
		logError(err, "Error al conectar a la base de datos")
		return false
	}
	err = DB.Ping()
	if err != nil {
		logError(err, "Error al verificar la conexión")
		return false
	}
	fmt.Println("Conexión exitosa a la base de datos.")
	return true
}

func GetDB() *sql.DB {
	return DB
}

func RunMigrations() error {
	driver, err := mysql.WithInstance(DB, &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"mysql", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func logError(err error, message string) {
	errorMessage := utils.ErrorServerMessage{
		Code:      400,
		Status:    "error",
		Error:     true,
		Message:   message,
		ErroType:  err.Error(),
		Timestamp: time.Now().Format(time.RFC3339),
	}
	jsonErr, _ := json.MarshalIndent(errorMessage, "", "  ")
	fmt.Println(string(jsonErr))
}
