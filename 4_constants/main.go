package main

import "fmt"

func main() {
	// single consts
	const name = "golang"
	// name = "hi"
	fmt.Println(name)

	// multipul const or const grouping
	const (
		port   = 5022
		pi     = 3.1416
		myName = "shahriar"
	)
	// fmt.Println(port)
	// fmt.Println(pi)
	// fmt.Println(myName)
	fmt.Println(port, pi, myName)
}
