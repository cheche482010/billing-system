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
}

type ResponseValidation struct {
	Status bool
	Error  int64
}

func ValidateData(data interface{}, validations map[string]ValidationCriteria) (interface{}, error) {
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
			errorMessage = ErrorValidationMessage{
				Code:      "400",
				Status:    "error",
				Error:     true,
				Message:   "Los datos enviados no son válidos.",
				Errors:    errorMessages,
				Timestamp: time.Now().Format(time.RFC3339),
			}
			return errorMessage, nil
		}
		return true, nil
	default:
		return false, nil
	}
}

func validatePattern(s string, pattern string) bool {
	match, _ := regexp.MatchString(pattern, s)
	return match
}
