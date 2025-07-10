package main

import "fmt"

func variadic(numbers ...int) { //  ... int mane multiple argument parameter hisabe pass hobe and seta se store korbe
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
	
}

func main() {
	variadic(5, 5, 6, 8, 1, 5) // that means ami function e multiple argument pass kortesi eta ke bole variadic function

}
