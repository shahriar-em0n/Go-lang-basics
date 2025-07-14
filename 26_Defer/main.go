package main

import "fmt"

func calculate() (reseult int) {
	
	fmt.Println("first - ", reseult)

	show := func() {
		reseult = reseult + 10
		fmt.Println("defer - ", reseult)
	}

	defer show()

	reseult = 5

	p := func (a int)  {
		fmt.Println("iam GPA - ", a)
	}

	defer p(reseult)

	defer fmt.Println(reseult)

	fmt.Println("second - ", reseult)
	return
}

// func calc() int {
	
// 	result := 0

// 	fmt.Println("first - ", result)

// 	show := func() {
// 		result = result + 10
// 	}

// 	defer show()

// 	result = 5
// 	fmt.Println("second - ", result)

// 	return result
// }

func main() {
	a := calculate()
	fmt.Println("first main - ",a)

	// b := calc()
	// fmt.Println("second main - ",b)

}


/* 2 Phase

1. compilationk phase
2. execution phase

### compilation phase

calculate = func(){...}
calculate anonymous  = func(){...}
calc = func(){...}
calc anonymous = func(){...}
main = func(){...}
---------------

Go have 2 rules when using defer in named return values & without named return values

	1. Named return values
		
		i. All codes execute
		ii. Defer function store kora hobe magic box e
		iii. Return -> all defer functions execute korbe
		iv. Return korbe named variables gular values

	2.Just return types

		i. All codes execute
		ii. Defer function store kora hobe magic box e
		iii. Returns values are evaluated at this time (store the return value)
		iv.All defer function execute


 

*/



/* Function list

	1. Defer function - last in first out
	2. parameter Vs Argument
	3. First order Function
	4. Init function - you can not call this, computer calls this automaticcally
	5. Standard or named function
	6.Anonymous function
	7.IIFE - Immediately invoked function expression
	8. Function expression or assign function in variable 
	9. Higher order function or first calss function(Threated as First class citizen)
	10.Callback function
	11. Receiver function
	12. Clousre-close over
	13. Variadic function
		i. Array
		ii. Slice
		iii. Pointer

*/