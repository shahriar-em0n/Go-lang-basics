package handlers

import (
	"eccomarce/database"
	"eccomarce/util"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pId, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "please give me valid product id", 400)
		return
	}

	for _, product := range database.Carlists {
		
		if product.Id == pId {
			util.SendData(w, product, 200)
			return
		}
	}

	util.SendData(w, "data pai nai", 404)

}
