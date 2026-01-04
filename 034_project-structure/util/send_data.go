package util

import (
	"encoding/json"
	"net/http"
)


func SendData(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}