package httphelpers

import (
	"encoding/json"
	"log"
	"net/http"
)

// ApiAnswer is a standard answer for an API
type ApiAnswer struct {
	// Status is the HTTP status code
	Status int `json:"status,omitempty"`

	// Message is a human-readable message
	Message string `json:"message,omitempty"`

	// Data is the data to be returned (if needed)
	Data any `json:"data,omitempty"`

	// Error is the error code (if any)
	Error string `json:"error,omitempty"`
}

// WriteResponse is a helper function to write a response to the response writer
func WriteResponse(status int, message string, data any, error string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		apiAnswer := ApiAnswer{
			Status:  status,
			Message: message,
			Data:    data,
			Error:   error,
		}

		apiAnswerJson, err := json.Marshal(apiAnswer)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"status":500,"message":"Internal Server Error","error":"internal_server_error"}`))
			return
		}

		w.WriteHeader(status)

		if _, err = w.Write(apiAnswerJson); err != nil {
			log.Printf("Error writing response: %s", err)
		}
	}
}

// WriteSuccess is a helper function to write a successful answer to the response writer
func WriteSuccess(code int, message string, data any) func(w http.ResponseWriter, r *http.Request) {
	return WriteResponse(code, message, data, "")
}

// WriteError is a helper function to write an error to the response writer
func WriteError(code int, error string, message string) func(w http.ResponseWriter, r *http.Request) {
	return WriteResponse(code, message, nil, error)
}
