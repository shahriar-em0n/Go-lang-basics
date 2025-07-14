package main

import "fmt"

var a = 10

func main() {
	age := 20

	if age >= 18 {
		a := 20

		fmt.Println(a)
	}
	fmt.Println(a)
}
