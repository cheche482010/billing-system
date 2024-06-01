package utils

import (
	"regexp"
	"strconv"
	"time"
)

type ValidationCriteria struct {
	Require   bool
	Pattern   string
	MinLength int
	MaxLength int
	IsInteger bool
}

type ResponseValidation struct {
	Status bool
	Error  ErrorValidationMessage
}

func ValidateData(data interface{}, validations map[string]ValidationCriteria) ResponseValidation {
	var errorMessage ErrorValidationMessage
	switch v := data.(type) {
	case string:
		errorMessages := make(map[string][]string)
		for field, criteria := range validations {
			if criteria.Require && len(v) == 0 {
				errorMessages[field] = append(errorMessages[field], field+" es obligatorio.")
			}
			if criteria.Pattern != "" {
				if !validatePattern(v, criteria.Pattern) {
					errorMessages[field] = append(errorMessages[field], field+" debe cumplir con el patrón especificado.")
				}
			}
			if criteria.MinLength != 0 && len(v) < criteria.MinLength {
				errorMessages[field] = append(errorMessages[field], field+" debe tener al menos "+strconv.Itoa(criteria.MinLength)+" caracteres.")
			}
			if criteria.MaxLength != 0 && len(v) > criteria.MaxLength {
				errorMessages[field] = append(errorMessages[field], field+" no debe exceder "+strconv.Itoa(criteria.MaxLength)+" caracteres.")
			}
		}
		if len(errorMessages) > 0 {
			errorMessage = ErrorMessage("Los datos enviados no son válidos.", errorMessages)
			return ResponseValidation{Status: false, Error: errorMessage}
		} else {
			return ResponseValidation{Status: true}
		}
	case int:
		errorMessages := make(map[string][]string)
		for field, criteria := range validations {
			if criteria.Require && v == 0 {
				errorMessages[field] = append(errorMessages[field], field+" es obligatorio.")
			}
			if criteria.IsInteger && v != int(v) {
				errorMessages[field] = append(errorMessages[field], field+" debe ser un número entero.")
			}
		}
		if len(errorMessages) > 0 {
			errorMessage = ErrorMessage("Los datos enviados no son válidos.", errorMessages)
			return ResponseValidation{Status: false, Error: errorMessage}
		} else {
			return ResponseValidation{Status: true}
		}
	default:
		errorMessage = ErrorMessage("El dato enviado no es válido o no se pudo procesar.", map[string][]string{"dato": {"Tipo de dato no soportado."}})
		return ResponseValidation{Status: false, Error: errorMessage}
	}
}

func validatePattern(s string, pattern string) bool {
	match, _ := regexp.MatchString(pattern, s)
	return match
}

func ErrorMessage(messageMain string, errorsMap map[string][]string) ErrorValidationMessage {
	return ErrorValidationMessage{
		Code:      "400",
		Status:    "error",
		Error:     true,
		Message:   messageMain,
		Errors:    errorsMap,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
