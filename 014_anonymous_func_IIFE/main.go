package main

import "fmt"

func add(x, y int) {
	fmt.Println(x + y)
}

func main() {
	var (
		a = 10
		b = 20
	)
	add(a, b) // we ivoke (invoke mane call kora)

	var value = 10
	if value > 0 {
		fmt.Println("value is greater than 0")
	}
	// anonymous function

	func(d, c int) {
		fmt.Println(d + c)
	}(a, b)
}

func init(){
	fmt.Println("hey i coder how are you")
}

// package main

// import "fmt"
// /*func add -----> ke standard or named function bole*/
// func add(a , b int){
// 	fmt.Println(a + b)
// }

// func main() {
// 	d := 5
// 	e := 5
// 	// we invoke ( invoke mane /call kora ) the add function here
// 	add(d,e)

// 	value := 10 // expression
// 	// if expression
// 	if value>0 {// if block
// 		fmt.Println("Value is greater than 0")
// 	}

// 	// anonymouse function
// 	// immediately invoked function expression
// 	func (a , b int)  {
// 		c := a + b
// 		fmt.Println(c)
// 	}(d , e)
// }

// func init(){
// 	fmt.Println("hey i`m the first function who can executed!")
// }
