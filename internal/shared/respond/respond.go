package respond

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := Response{
		Status:  true,
		Message: message,
		Data:    data,
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func Error(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := Response{
		Status:  false,
		Message: message,
		Data:    data,
	}

	_ = json.NewEncoder(w).Encode(resp)
}