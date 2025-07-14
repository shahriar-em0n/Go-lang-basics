package main

import "fmt"

func Welcome() {
	var name string
	fmt.Print("Enter your Name - ")
	fmt.Scanln(&name)
	fmt.Println("Welcome", name)
}

func buyCar() {
	var condi string
	fmt.Print("Can you buy a car? (Yes/No) - ")
	fmt.Scanln(&condi)

	if condi == "Yes" {
		carBrand, carModel, carColor, carType := cardetails()
		display(carBrand, carModel, carColor, carType)
	} else {
		fmt.Println("Have nice day")
	}
}

func cardetails() (string, string, string, string) {
	var carBrand, carModel, carColor, carType string

	fmt.Print("Enter the Car Brand Name - ")
	fmt.Scanln(&carBrand)

	fmt.Print("Enter the Car Model - ")
	fmt.Scanln(&carModel)

	fmt.Print("Enter the Car Color - ")
	fmt.Scanln(&carColor)

	fmt.Print("Enter the Car Type (Electric or Fuel) - ")
	fmt.Scanln(&carType)

	return carBrand, carModel, carColor, carType
}

func display(carBrand, carModel, carColor, carType string) {
	fmt.Println("\nThis is your Cash Memo")
	fmt.Println("Car Brand  -", carBrand)
	fmt.Println("Car Model  -", carModel)
	fmt.Println("Car Color  -", carColor)
	fmt.Println("Car Type   -", carType)
	fmt.Println("Total cost is - 1cr")
	fmt.Println("Thank you for your order Have a nice day")
}

func main() {
	Welcome()
	buyCar()
}
