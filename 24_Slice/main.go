package main

import (
	"fmt"
)

func changeSlice(a []int) []int {
	a[0] = 10
	b := append(a /* Header */, 11)
	return b

}

func main() {

	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7)

	a := x[4:]

	fmt.Println(a)

	y := changeSlice(a)


	fmt.Println("value of x : ", x[0])
	fmt.Println("value of y : ", y[0])

	// fmt.Println(&x[0:8][0])

}


/*

	what i learn :


	1. Slice from an existing array
	2. Slice from an slice
	3. Slice literal
	4. Slice Structure
		i. Pointer
		ii. Lenght
		iii. Capacity
	5. Slice in make fucntion with  len & cap
	6. Empty slice or nil slice
	7. Slice underlying array rule ==> 1024 --> 100% increse that means 1024 * 2 times but 1024 theke boro hole 25% increase hobe

	----------------

	*** Code segment ***

	main = func(){...}

*/
