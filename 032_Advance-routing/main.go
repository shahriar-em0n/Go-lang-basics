package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /wlc", http.HandlerFunc(welcome))

	mux.Handle("GET /cars",  (http.HandlerFunc(getcars)))

	mux.Handle("POST /addCar",  (http.HandlerFunc(addCars)))

	fmt.Println("Server running on :8080")

	globalRouter := globalRouter(mux)

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Server error starting", err)
	}
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
