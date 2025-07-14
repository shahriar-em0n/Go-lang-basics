package main

import "fmt"

func main() {
	// they are many ways to declear variables

	var name string ="hellow"
	var age int = 18
	var isAdult bool = true
	var price float32 = 50.5
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isAdult)
	fmt.Println(price)


	//2nd type => infer 
	var name2 = "golang"
	var price2 = 56.02
	fmt.Println("this is infer declearation =>",name2)
	fmt.Println(price2)

	//3rd type => shorthand syntax
	
	name3 := "hello shahriar this is you captain"
	fmt.Println(name3)

	price3 := 8456
	fmt.Println(price3)
	





}
