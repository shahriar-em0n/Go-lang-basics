package main

import "fmt"

var (
	a = 20
	b = 30
)

func fprint(num int) {
	fmt.Println(num)
}

/*   --------------------> egula sob gulai standerd function           */

func add(n int, y int) {
	sum := n + y
	fprint(sum)
}

func main() {
	add(a, b)

}
