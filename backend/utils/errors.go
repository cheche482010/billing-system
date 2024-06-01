package utils

type ErrorServerMessage struct {
	Code      int    `json:"code"`
	Status    string `json:"status"`
	Error     bool   `json:"error"`
	Message   string `json:"message"`
	ErroType  string `json:"errotype"`
	Timestamp string `json:"Timestamp"`
}

type ErrorRouteMessage struct {
	Code      int    `json:"code"`
	Status    string `json:"status"`
	Error     bool   `json:"error"`
	Message   string `json:"message"`
	ErrorType string `json:"errorType"`
	Path      string `json:"path"`
	Timestamp string `json:"timestamp"`
}

type ErrorSqlMessage struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	ErroType  string `json:"erroType"`
	Timestamp string `json:"timestamp"`
}

type ErrorValidationMessage struct {
	Code      string              `json:"code"`
	Status    string              `json:"status"`
	Error     bool                `json:"error"`
	Message   string              `json:"message"`
	Errors    map[string][]string `json:"errors"`
	Timestamp string              `json:"timestamp"`
}

type ErrorMethodMessage struct {
	Code      int    `json:"code"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	ErrorType string `json:"errortype"`
	Timestamp string `json:"timestamp"`
}

type ErrorJSONdMessage struct {
	Code      int    `json:"code"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	ErrorType string `json:"errortype"`
	Timestamp string `json:"timestamp"`
}
