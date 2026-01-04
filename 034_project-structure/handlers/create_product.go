package handlers

import (
	"ecom/database"
	"ecom/util"
	"encoding/json"
	"fmt"
	"net/http"
)


func RequestCar(w http.ResponseWriter, r *http.Request) {
	var newCar database.Car

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newCar)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me a valid json", 400)
		return
	}

	newCar.ID = len(database.CarList) + 1
	 database.CarList = append(database.CarList, newCar)

	util.SendData(w, newCar, 201)
}