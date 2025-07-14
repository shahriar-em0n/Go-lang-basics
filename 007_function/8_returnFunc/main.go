package main

import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2

	return sum
}

func numbers(num1 int, num2 int) (int, int, int, int) {
	sum := num1 + num2 // 30

	sub := num1 - num2 // 10

	mult := num1 * num2 // 200

	divi := num1 / num2 // 2

	return sum, sub, mult, divi

}

func main() {
	a := 20
	b := 10

	// x , n := numbers(a,b)
	x, n, z, y := numbers(a, b)
	fmt.Println("sum  = ", x)
	fmt.Println("sub  = ", n)
	fmt.Println("mult = ", z)
	fmt.Println("divi = ", y)

}
