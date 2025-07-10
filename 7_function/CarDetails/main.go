package main

import "fmt"

func carDetails(Brand string, Model string, colure string, Type string, Gener string, Engine string) {
	fmt.Println("CarBrand 	= ", Brand)
	fmt.Println("CarMoel 	= ", Model)
	fmt.Println("CarColore 	= ", colure)
	fmt.Println("CarGener 	= ", Gener)
	fmt.Println("CarType 	= ", Type)
	fmt.Println("CarEngine 	= ", Engine)
}

func main() {
	carBrand := "BMW"
	carModel := "BMW I 8"
	carColure := "Red"
	carType := "fuel"
	carGener := "sports"
	carEngine := "2000 HP"

	carDetails(carBrand, carModel, carColure, carGener, carType, carEngine)
}
