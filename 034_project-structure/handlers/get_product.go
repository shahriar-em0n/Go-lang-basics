package handlers

import (
	"ecom/database"
	"ecom/util"
	// "ecom/util"
	// "log"
	"net/http"
	"strconv"
)

func GetCar(w http.ResponseWriter, r *http.Request) {
	// util.SendData(w, database.CarList, 200)

	carId := r.PathValue("getCarId")

	cId, err := strconv.Atoi(carId)
	if err != nil {
		http.Error(w, "Please give me a valid car id", 400)
		return
	}

	for _, car := range database.CarList {
		if car.ID == cId {
			util.SendData(w, car, 200)
			return
		}
	}

	util.SendData(w, "data pai nai", 404)
}
