package response

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func write(w http.ResponseWriter, status int, success bool, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(APIResponse{
		Success: success,
		Message: message,
		Data:    data,
	})
}

func OK(w http.ResponseWriter, data interface{}) {
	write(w, http.StatusOK, true, "", data)
}

func Created(w http.ResponseWriter, message string, data interface{}) {
	write(w, http.StatusCreated, true, message, data)
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func BadRequest(w http.ResponseWriter, message string) {
	write(w, http.StatusBadRequest, false, message, nil)
}

func Unauthorized(w http.ResponseWriter, message string) {
	write(w, http.StatusUnauthorized, false, message, nil)
}

func NotFound(w http.ResponseWriter, message string) {
	write(w, http.StatusNotFound, false, message, nil)
}

func InternalServerError(w http.ResponseWriter, message string) {
	write(w, http.StatusInternalServerError, false, message, nil)
}
