package main

import "fmt"

/* Higher order Function */
func processOperation(a int, b int, Op func(p int, q int)/*This function called callbacke function*/){
	Op(a, b)
}

/* Higher order Return function */

func call() func(s int, d int){
	return add
}

func add(x int, y int) { // parameter
	sum := x + y
	fmt.Println(sum)
}

func main() {
	add(4, 5) // argument

	processOperation(2, 5, add)

	sum := call() // sum is a function expression
	sum(5, 5)
}


/* 1. Parameter & argument
   2. First order function
	  
	  i. Standard function or named fuction
	  ii. Anonymous function
	  iii. IIFE
	  iv. Function expression

	3. Higer order function 

	   i. If any function called Higher order function than any one of the following 3 
	   		1. Parameter function hisabe asbe
			2. Function return korbe
			3. Both

	   ii. First order and Higer order function are  mainly Functional paradigm Langugae Example -> Haskel, racket, etc

	   iii. Functional paradigm Languge are inspierd to Discrete mathematics --> LOgic (chapter)

	   	   i. First order logic
		   ii. Higher order logic

		### Logic ###

		1. Object (People, animal, car, etc)
		2. Property (Color, student)
		3. Relation
		Example :
			i. Shahriar(object & Property) is a student
			ii. Apple is red(property)
		
		### First order Logic

		are mainly work (Object, Property, Relation)
		 
		Rule : All customers(Object) must pay their  pizza bills.
			   
		       All students must wear their uniforms.
		
		### Higher order logic

		are mainly work (Object, Property, Relation) and Rules 

		i. Rule : Any rule that applies to all customers must also apply to Shahriar

		ii Rule : All customers must pay tips to the waiter 
	
	4. Callback function
	
       example :  func processOperation(a int, b int, Op func(p int, q int)/*This function called callbacke function){
	   Op(a, b) 

	5. First class citizen ==> ( variable assignee data )

	   Example : a := int, float, string, bool

*/