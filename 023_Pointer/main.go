package main

import "fmt"

func main() {
	x := 20

	fmt.Println("before using pointer x value : ", x)

	p := &x

	*p = 30

	fmt.Println("Address : ", &p) // & -- Print the address of P - 0xc00005a060

	fmt.Println("Address : ", p) // & -- Print the address of x - 0xc00000a0e8
	//fmt.Println("Value at address : ", *p) //* -- print the address value

	fmt.Println("After using pointer x value : ", x)
}

// package main

// import "fmt"

// // func print(numbers [3] int){ // pass by value -- without pointer
// // 	fmt.Println(numbers)
// // }

// func print(numbers *[3] int){ // pass by reference -- with reference
// 	fmt.Println(numbers)
// }

// func main() {
// 	x := 20 // pointer is a address of memory

// 	address := &x // & --> Ampersand

// 	*address = 30

// 	fmt.Println("x = ", address)

// 	fmt.Println("Address of x : ",address)
// 	fmt.Println("Value of address : ",*address)

// 	arr := [3] int{1, 2, 3}
// 	print(&arr) // pass by reference

// 	// arr := [3] int{1, 2, 3}
// 	// print(arr) // pass by value
// }
