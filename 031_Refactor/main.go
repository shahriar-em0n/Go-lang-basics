package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	carlists []Cars
)

type Cars struct {
	Id       int     `json:"id"`
	Brand    string  `json:"brand"`
	Model    string  `json:"model"`
	Color    string  `json:"color"`
	Price    float32 `json:"price"`
	ImageUrl string  `json:"imageUrl"`
}



func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to our car showroom sir")
	fmt.Fprintln(w, "please enter the route for showing our car /cars")
}

func addCars(w http.ResponseWriter, r *http.Request) {
	// solve the CORS error

	handleCors(w)
	handlePreflightReq(w, r)

	if r.Method != "POST" {
		http.Error(w, "Please give me valid POST request", 400)
	}

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

func getcars(w http.ResponseWriter, r *http.Request) {
	// solve the CORS problem
	handleCors(w)
	if r.Method != http.MethodGet {
		http.Error(w, "Please give me a valid GET request", 400)
		return
	}

	sendData(w, carlists, 200)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/wlc", welcome)
	mux.HandleFunc("/cars", getcars)
	mux.HandleFunc("/addCar", addCars)

	fmt.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Server error starting", err)
	}
}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func handlePreflightReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int){
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func init() {
	car1 := Cars{
		Id:       1,
		Brand:    "TOYOTA",
		Model:    "Crown",
		Color:    "Red",
		Price:    55000,
		ImageUrl: "https://cdn.motor1.com/images/mgl/JvQ8Q/s3/2021-toyota-crown-update.jpg",
	}

	car2 := Cars{
		Id:       2,
		Brand:    "BMW",
		Model:    "M5",
		Color:    "Blue",
		Price:    102000,
		ImageUrl: "https://i.pinimg.com/736x/7c/06/4c/7c064c3ed05fe1fed6367f5c6765aa0e.jpg",
	}

	car3 := Cars{
		Id:       3,
		Brand:    "MERCEDES",
		Model:    "S-Class",
		Color:    "Black",
		Price:    110000,
		ImageUrl: "https://i.pinimg.com/736x/a0/62/66/a06266d4d30bdd45468516feea72e7fa.jpg",
	}

	car4 := Cars{
		Id:       4,
		Brand:    "TESLA",
		Model:    "Model S",
		Color:    "White",
		Price:    95000,
		ImageUrl: "https://i.pinimg.com/736x/ee/82/f0/ee82f034de6940131f6e4de8f49049ad.jpg",
	}

	car5 := Cars{
		Id:       5,
		Brand:    "AUDI",
		Model:    "A8",
		Color:    "Gray",
		Price:    86000,
		ImageUrl: "https://i.pinimg.com/736x/d5/e4/7f/d5e47f2b101d4cc11ab7a79a094cc904.jpg",
	}

	car6 := Cars{
		Id:       6,
		Brand:    "FORD",
		Model:    "Mustang GT",
		Color:    "Yellow",
		Price:    72000,
		ImageUrl: "https://i.pinimg.com/1200x/ce/9c/3b/ce9c3bb84bc7674a32b6b1f9b5af9a5a.jpg",
	}

	carlists = append(carlists, car1)
	carlists = append(carlists, car2)
	carlists = append(carlists, car3)
	carlists = append(carlists, car4)
	carlists = append(carlists, car5)
	carlists = append(carlists, car6)
}
