package main

import (
	"fmt"
)

type student struct {
	Name       string
	Department string
	session    int
	Roll       int
	Regi       int
	BloadGroup string
}

func (studentId student) printStudentDetails() { // Recevier function
	fmt.Println("----Student Details-----")
	fmt.Println("Name : ", studentId.Name)
	fmt.Println("Deparment : ", studentId.Department)
	fmt.Println("Session : ", studentId.session)
	fmt.Println("Roll : ", studentId.Roll)
	fmt.Println("Registration : ", studentId.Regi)
	fmt.Println("Bload Group : ", studentId.BloadGroup)
}

func Shahriar() {
	Shahriar := student{
		Name:       "Mohammad Shahriar",
		Department: "CST",
		session:    2023,
		Roll:       880485,
		Regi:       12345897,
		BloadGroup: "AB+",
	}

	Shahriar.printStudentDetails()
}

func Fahim() {
	Fahim := student{
		Name:       "Fahim Hossain",
		Department: "CST",
		session:    2023,
		Roll:       88658,
		Regi:       12345577,
		BloadGroup: "AB+",
	}

	Fahim.printStudentDetails()
}

func main() {
	Shahriar()
	Fahim()
}

// package main

// import "fmt"

// type Person struct {
// 	Name       string
// 	Hight      string
// 	Age        int
// 	BloadGroup string
// }

// // Recevier function

// func (perso Person) printPersonDetails() {
// 	fmt.Println("Hello Uesr")
// 	fmt.Println("Name : ", perso.Name)
// 	fmt.Println("Age : ", perso.Age)
// 	fmt.Println("Hight : ", perso.Hight)
// 	fmt.Println("Bload Group : ", perso.BloadGroup)
// }

// // func printPersonDetails(perso Person) {
// // 	fmt.Println("Hello Uesr")
// // 	fmt.Println("Name : ", perso.Name)
// // 	fmt.Println("Age : ", perso.Age)
// // 	fmt.Println("Hight : ", perso.Hight)
// // 	fmt.Println("Bload Group : ", perso.BloadGroup)
// //}

// type CarBrand struct {
// 	brandName string
// 	carModel  string
// 	madeIn    string
// }

// func (car CarBrand) PrintCarDetails() {

// 	fmt.Println("Hello Uesr")
// 	fmt.Println("Brand Name : ", car.brandName)
// 	fmt.Println("Car Model : ", car.carModel)
// 	fmt.Println("Made In : ", car.madeIn)
// }

// func main() {
// 	shahriar := Person{
// 		Name:       "Mohammad Shahriar",
// 		Hight:      "5`5",
// 		Age:        18,
// 		BloadGroup: "A+",
// 	}
// 	shahriar.printPersonDetails()
// 	//printPersonDetails(shahriar)

// 	BMW := CarBrand{
// 		brandName: "BMW",
// 		carModel:  "I8",
// 		madeIn:    "Germany",
// 	}

// 	BMW.PrintCarDetails()
// 	// PrintCarDetails(BMW)

// 	Rohim := Person{
// 		Name:       "Rohim",
// 		Hight:      "5`8",
// 		Age:        20,
// 		BloadGroup: "AB+",
// 	}
// 	Rohim.printPersonDetails()
// }
