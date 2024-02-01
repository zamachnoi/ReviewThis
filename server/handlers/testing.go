package handlers

// create a handler that will return a 200 response with a JSON body containing the message "Hello from testing handler"

import (
	"encoding/json"
	"net/http"
)

type TestingResponse struct {
	Message string `json:"message"`
}

func TestingHandler(w http.ResponseWriter, r *http.Request) {
	response := TestingResponse{
		Message: "Hello from testing handler",
	}
	json.NewEncoder(w).Encode(response)
}