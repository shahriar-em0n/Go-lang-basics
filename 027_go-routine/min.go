package main

import (
	"fmt"
	"time"
)

var a = 10
const p = 11

func printHello(num int){
	fmt.Println("Hello shahriar", num)
}

func main(){

	fmt.Println("hello go routin")
	go printHello(1) 

	go printHello(2)

	go printHello(3)

	go printHello(4)

	go printHello(5)

	fmt.Println(a, " ", p)

	time.Sleep(1 * time.Second)
	 
} 