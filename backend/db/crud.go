package db

import (
	"database/sql"
	"fmt"
	"log"
)

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Error   bool        `json:"error"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Insert(sqlQuery string, params map[string]interface{}) error {

	paramsList := make([]interface{}, len(params))
	i := 0
	for _, paramValue := range params {
		paramsList[i] = paramValue
		i++
	}

	result, err := DB.Exec(sqlQuery, paramsList...)
	if err != nil {
		log.Printf("Error al insertar: %v\n", err)
		return err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error obteniendo el ID del último registro insertado: %v\n", err)
		return err
	}
	fmt.Printf("Registro insertado con ID: %d\n", lastID)
	return nil
}

func Update(sqlQuery string, params map[string]interface{}) error {

	paramsList := make([]interface{}, len(params))
	i := 0
	for _, paramValue := range params {
		paramsList[i] = paramValue
		i++
	}

	result, err := DB.Exec(sqlQuery, paramsList...)
	if err != nil {
		log.Printf("Error al actualizar: %v\n", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error obteniendo el número de filas afectadas: %v\n", err)
		return err
	}
	fmt.Printf("Registros actualizados: %d\n", rowsAffected)
	return nil
}

func Delete(sqlQuery string, params map[string]interface{}) error {

	paramsList := make([]interface{}, len(params))
	i := 0
	for _, paramValue := range params {
		paramsList[i] = paramValue
		i++
	}

	result, err := DB.Exec(sqlQuery, paramsList...)
	if err != nil {
		log.Printf("Error al eliminar: %v\n", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error obteniendo el número de filas eliminadas: %v\n", err)
		return err
	}
	fmt.Printf("Registros eliminados: %d\n", rowsAffected)
	return nil
}

func GetBy(sqlQuery string, params map[string]interface{}) (*sql.Row, error) {
	// Convertir el mapa de parámetros a una lista de interfaces
	paramsList := make([]interface{}, len(params))
	i := 0
	for _, paramValue := range params {
		paramsList[i] = paramValue
		i++
	}

	row := DB.QueryRow(sqlQuery, paramsList...)
	return row, nil
}
