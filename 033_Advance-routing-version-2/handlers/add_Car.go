package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"eccomarce/database"
	"eccomarce/util"
)

func AddCars(w http.ResponseWriter, r *http.Request){

	var newCars database.Cars

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newCars)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid Json", 400)
		return
	}

	newCars.Id = len(database.Carlists) + 1

	database.Carlists = append(database.Carlists, newCars)

	util.SendData(w, newCars, 201)

}
