package main

import "fmt"

var(
	a = 20 
	b = 40
)


func sum(x int, y int) {
	z := x + a
	fmt.Println(z)
}

func main() {

	c := 20
	// d := 20

	sum(c,a)

}