package main

import "fmt"

func main() {
	// var arr [2]int
	// arr[0]= 3
	// arr[1]=6

	// arr := [2]int{5, 3}


	// country := [5]string{"Bangladesh", "Finland", "Paris", "Japan", "China"}

	// fmt.Println(arr)
	// fmt.Println(country)
	// fmt.Println(country[0])
	// fmt.Println(country[1])
	// fmt.Println(country[2])
	// fmt.Println(country[3])
	// fmt.Println(country[4])

	// array using for loop

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(country[i])
	// }

	// array input using loop

	var studnetName [5] string

	for i := 0 ; i < 5 ; i++{
		fmt.Print("Enter Students name : ")
		fmt.Scan(&studnetName[i])
	}

	fmt.Println(studnetName)

}
