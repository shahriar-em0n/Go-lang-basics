package main

import (
	"net/http"
)

func getcars(w http.ResponseWriter, r *http.Request) {
	sendData(w, carlists, 200)
}