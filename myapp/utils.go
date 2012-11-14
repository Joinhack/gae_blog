package myapp

import (
	"net/http"

	"encoding/json"
)

func OutputJson(w http.ResponseWriter, data interface{}) (err error) {
	w.Header().Set("content-type", "text/plain; charset=UTF-8")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(data)
	return
}