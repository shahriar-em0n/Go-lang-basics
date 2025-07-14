package main

import "fmt"

func StudentDetails(School string, Name string, Roll int, Subject string, Sex string, Age int, Number string) {
	fmt.Println("School Name 	= ", School)
	fmt.Println("Students Name 	= ", Name)
	fmt.Println("Student Roll 	= ", Roll)
	fmt.Println("Subjact		= ", Subject)
	fmt.Println("Sex 		= ", Sex)
	fmt.Println("Age 		= ", Age)
	fmt.Println("Phone Number 	= ", Number)
}

func main() {
	SchoolName := "BEPZA PUBLIC SCHOOL AND COLLAGE"
	StudentName := "Mohammad Shahriar"
	StudentRoll := 1528
	subjact := "Science"
	Sex := "MALE"
	Age := 19
	PhoneNUmber := "01782291277"

	StudentDetails(SchoolName, StudentName, StudentRoll, subjact, Sex, Age, PhoneNUmber)
}
