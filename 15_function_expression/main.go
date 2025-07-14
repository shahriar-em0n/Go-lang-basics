package main

import "fmt"
/*func add -----> ke standard or named function bole*/
// func add(a , b int){
// 	fmt.Println(a + b)
// }

func main() {

	// anonymouse function
	// immediately invoked function expression
	add := func(a int , b int)  {
		c := a + b
		fmt.Println(c)
	}
	add(4 , 7)
}

func init(){
	fmt.Println("hey i`m the first function who can executed!")
}
