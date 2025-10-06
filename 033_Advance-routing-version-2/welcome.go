package main

import (
	"fmt"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to our car showroom sir")
	fmt.Fprintln(w, "please enter the route for showing our car /cars")
}