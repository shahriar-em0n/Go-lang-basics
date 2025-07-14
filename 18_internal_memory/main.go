package main

import "fmt"

const a = 40

var (
	p = 20
)

func add(x int, y int){
	sum := x + y
	fmt.Println(sum)
}//Class number --> 24

func call() {
	add := func(x int, y int) {
		sum := x + y
		fmt.Println(sum)
	}

	add(5, 5)
}

func main() {
	 add(5, 4)
	 add(a, 4)// Class number --> 24
	call()
	b := 30
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(p)
}

func init() {
	fmt.Println("hello world!")
}


/*
	compile time :
		a = 40
		add = func(){...}
		call = func(){..}
		main = func(){..}
		init = func(){..}

	Run time :

		

*/

/*
	There is a 2 Phase :
	1. Compilation phase
		***Code Segment ***

		a = 40
		call = func(){...}
		add = func(){...}
		main = func(){...}
		init = func(){...}


	2. Execution phase

	i. jokhon ami go run use kori tokhon computer seta compile kore and sathe sathe seta run kore ./main er maddhome
		go run main.go  ==> compile it ==> main ==> ./main.go

	ii. jokhon ami go build use kori tokhon computer seta compile kore but seta run kore na
		go build main.go ==> compile it ==> main 
		eta amra nijerai execute korte pari ./main use kore

*/

/* 
	***Binary File ***

	after compile this is main file

	binary file er bhitore code thake sudu 011001001110101 line by line amon binary code thakbe

	main file :
	0011101011010010
	0011101011010010
	0011101011010010
	0011101011010010
	0011101011010010


	ami jodi ei binary file ke build korte chai tahole amake 

	go build 18_internal_memory/main.go 

	ei command ta use korte hobe and code run korar jonno 

	./ main

	use korte hobe and eta joto bar ami code build korbo mane code changes anbo totobar amake eti korte hobe

*/