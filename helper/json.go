package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {
	err := json.NewDecoder(r.Body).Decode(result)
	PanicIfError(err)
}

func WriteResponseBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	PanicIfError(err)
}
