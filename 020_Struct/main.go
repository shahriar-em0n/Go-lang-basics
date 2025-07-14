package main

import "fmt"

type Person struct { // member varibale or property
	Name       string
	Hight      string
	Age        int
	BloadGroup string
}


func printPersonDetails(perso Person) {
	fmt.Println("Hello Uesr")
	fmt.Println("Name : ", perso.Name)
	fmt.Println("Age : ", perso.Age)
	fmt.Println("Hight : ", perso.Hight)
	fmt.Println("Bload Group : ", perso.BloadGroup)
}

func main() {
	shahriar := Person{ // instantiate ----- or instanc or object
		Name:       "Mohammad Shahriar",
		Hight:      "5`5",
		Age:        18,
		BloadGroup: "A+",
	}
	printPersonDetails(shahriar)
	
}

/*
struct {..}
printprd {..}
main{ .. }

// package main

// import "fmt"

// type User struct {
// 	Name   string // member varibale or property
// 	Age    int
// 	contry string
// }

// func printUserDetails(usr User){
// 	fmt.Println(usr.Name)
// 	fmt.Println(usr.Age)
// 	fmt.Println(usr.contry)
// }

// type userDetails struct {
// 	Name      string
// 	Age       int
// 	work      string
// 	Naonality string
// }

// type carDetails struct {
// 	brandName string
// 	carModle  string
// 	color     string
// 	carType   string
// }

// func main() {

// 	user1 := User{ // instantiate ----- or instanc or object
// 		Name:   "shahriar",
// 		Age:    30,
// 		contry: "Bangladesh",
// 	}
// 	printUserDetails(user1)
// 	// fmt.Println("name : ", user1.Name)
// 	// fmt.Println("age : ", user1.Age)
// 	// fmt.Println("Country : ", user1.contry)

// 	/*-------------------------------------*/

// 	user2 := User{ // Instance or object
// 		Name: "dhummachale",
// 		Age:  15,
// 	}

// 	fmt.Println("name : ", user2.Name)
// 	fmt.Println("age : ", user2.Age)

// 	/*-------------------------------------*/

// 	shahriar := userDetails{
// 		Name:      "Shahriar",
// 		Age:       19,
// 		work:      "student",
// 		Naonality: "Bangladesh",
// 	}

// 	fmt.Println("Name : ", shahriar.Name)
// 	fmt.Println("Name : ", shahriar.Age)
// 	fmt.Println("Name : ", shahriar.work)
// 	fmt.Println("Name : ", shahriar.Naonality)

// 	/*-------------------------------------*/

// 	BMW := carDetails{
// 		brandName: "BMW",
// 		carModle: "I8",
// 		color: "Red",
// 		carType: "Fuel",
// 	}

// 	fmt.Println("Show me the car ")

// 	fmt.Println("Brand		: ",BMW.brandName)
// 	fmt.Println("Car Model	: ",BMW.carModle)
// 	fmt.Println("Color 		: ",BMW.color)
// 	fmt.Println("Car Type	: ",BMW.carType)

// }
