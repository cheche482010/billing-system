package db

import (
	"database/sql"
	"fmt"
	"time"

	"billing-system/utils"

	_ "github.com/go-sql-driver/mysql"
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
		errorMessage := utils.ErrorServerMessage{
			Code:      400,
			Status:    "error",
			Error:     true,
			Message:   "Error al conectar a la base de datos",
			ErroType:  err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		}
		jsonErr, _ := json.MarshalIndent(errorMessage, "", "  ")
		fmt.Println(string(jsonErr))
		return false
	}
	err = DB.Ping()
	if err != nil {
		errorMessage := utils.ErrorServerMessage{
			Code:      400,
			Status:    "error",
			Error:     true,
			Message:   "Error al verificar la conexión",
			ErroType:  err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		}
		jsonErr, _ := json.MarshalIndent(errorMessage, "", "  ")
		fmt.Println(string(jsonErr))
		return false
	}
	fmt.Println("Conexión exitosa a la base de datos.")
	return true
}

func GetDB() *sql.DB {
	return DB
}
