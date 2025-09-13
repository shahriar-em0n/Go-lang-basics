package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func addCars(w http.ResponseWriter, r *http.Request) {

	var newCars Cars

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newCars)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid Json", 400)
		return
	}

	newCars.Id = len(carlists) + 1

	carlists = append(carlists, newCars)

	sendData(w, newCars, 201)

}
