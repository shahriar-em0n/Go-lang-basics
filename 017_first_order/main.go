package main

import "fmt"

/*Higher order function example 
	if any function is called higher order than there is a 3 reiQirmenets must have 

	1. parameter -> function
	2. function return
	3. both
*/
	// parameter function
/*func processOperation(a int, b int, op func(x int , y int)) {  
	 op(a, b)
}
*/
// return function

func call() func (x int , y int){
	return add
}

func add(x int , y int){
	z := x + y
	fmt.Println(z)
}

func main() {
	sum := call()
	sum(4,7)
}

/*
	1. First order function {

		i. Standard function or named function
		ii. anonymouse function
		iii. IIFE
		iv. function expression 
	}

	2. Higher order function

	# they both are create by functional paradigm language 
		example ---> haskel, racket etc
	
	functional paradigm are inspierd by math
		math ---> chapter name : logic (discrete mathematics)
		1. first order logic 
		2. higher order logic

			### Logic ###
	# first order logic

	1. object (people, animal, car etc)
	2. Property (color, student)
	3. Relation

	i. Shahriar(object) is a student(property)
	ii. apple is red(property)
	iii. shahriar is smaller than rakib (hole thinks are relation)

	example --->
	Rule :  All customers must pay their pizza bills.

	Rule : All students must wear their uniforms


	### Higher order logic

	Higher order logics are basicaly work in Rules
example ---> 
	Any rule that applies to all customer must also apply to shahriar

example --->
	rule 1 : all customers must pay tips  waiter

	3. Call backe function

		what is callback function?
		higher order function e jodi kono function ke pass kora hoy setai callback function
	
	4. first class function or cityzen ---> 
		(varibale er bhitor ami je data gulo assigne kori segulakei first class cityzen bole)

*/