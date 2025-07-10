package main

import "fmt"

const a = 10

var p = 100

func outer() (func(), func()) {
	money := 100
	age := 30

	salary := 50000

	fmt.Println("age = ", age)

	show := func() {
		money = money + a + p
		fmt.Println(money)
	}

	myIncome := func() {
		salary = salary + money + a + p
		fmt.Println(salary)
	}

	return show, myIncome
}

func main() {
	fmt.Println("===Bank===")
	use1, user2 := outer() // show , myIncome
	use1()                 // show - 210
	use1()					

	user2() // myIncome

	user1, use2 := outer()
	use2()
	use2()

	user1()

}

/*	There is a 2 Phase :
	1. Compilation phase (compile time )
		***Code Segment ***

		a = 10
		outer = func(){...}
		outerAnnonymouse1 = func(){...}
		outerAnnonymouse2 = func(){...}
		main = func(){...}



*/

// package main

// import "fmt"

// const a = 10 // a is a constant thats why a get store code segment

// var p = 100

// func outer() func(){
// 	money := 100
// 	age := 30

// 	fmt.Println("age = ", age)

// 	show := func ()  { // basically show is a outer anonymous function
// 		money = money + a + p
// 		fmt.Println(money)
// 	}
// 	return show
// }

// func call(){
// 	incre := outer()
// 	incre()
// 	incre()

// 	incre2 := outer()
// 	incre2()
// 	incre2()
// }

// func main() {
// 	call()

// }

// func init(){
// 	fmt.Println("== Bank ==")
// }

// /*
// 	There is a 2 Phase :
// 	1. Compilation phase (compile time )
// 		***Code Segment ***

// 		a = 10
// 		outer = func(){...}
// 		outerAnonymous1 = func(){...}
// 		call = func(){...}
// 		main = func(){...}
// 		init = func(){...}

// 	2. Execution phase (run time)

// 	i. jokhon ami go run use kori tokhon computer seta compile kore and sathe sathe seta run kore ./main er maddhome
// 		go run main.go  ==> compile it ==> main ==> ./main.go

// 	ii. jokhon ami go build use kori tokhon computer seta compile kore but seta run kore na
// 		 ==> compile it ==> main
// 		eta amra nijerai execute korte pari ./main use kore

// */

// /*
// 	***Binary File ***

// 	after compile this is main file

// 	binary file er bhitore code thake sudu 011001001110101 line by line amon binary code thakbe

// 	main file :
// 	0011101011010010
// 	0011101011010010
// 	0011101011010010
// 	0011101011010010
// 	0011101011010010

// 	ami jodi ei binary file ke build korte chai tahole amake

// 	go build 18_internal_memory/main.go

// 	ei command ta use korte hobe and code run korar jonno

// 	./ main

// 	use korte hobe and eta joto bar ami code build korbo mane code changes anbo totobar amake eti korte hobe

// */
