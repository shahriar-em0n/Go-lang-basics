package handlers

import (
	"eccomarce/database"
	"eccomarce/util"
	"net/http"
)

func Getcars(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.Carlists, 200)
}