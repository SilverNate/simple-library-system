package utils

import (
	"encoding/json"
	"net/http"
)

// JsonResponse defines the standard JSON structure for responses
type JsonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"` // Optional field for data
}

type JsonPaginationResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Page    int    `json:"page"`
	Total   int    `json:"Total"`
	Data    any    `json:"data,omitempty"`
}

// SendJSONResponse sends a standard JSON response
func SendJSONResponse(w http.ResponseWriter, code int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(JsonResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendJSONResponseWithPagination(w http.ResponseWriter, code int, message string, data any, page, total int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(JsonPaginationResponse{
		Code:    code,
		Message: message,
		Data:    data,
		Page:    page,
		Total:   total,
	})
}
