package main

import "fmt"

var a = 20

func main() {
	fmt.Println("hello init function")
	fmt.Println(a)
}

func init(){
	fmt.Println("i am the first function that is executed first")
	fmt.Println(a)
}