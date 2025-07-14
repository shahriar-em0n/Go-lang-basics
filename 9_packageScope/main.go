package main

import (
	"fmt"

	"pkg/9_packageScope/mathlib"
	"pkg/9_packageScope/totalage"
)

var(
	a = 20
	b = 10
)

/*
1. go mod init example.com

2.If you can run single file than this command : go run 9_packageScope/main.go 9_packageScope/add.go --> ekhane tomar file er name diva like of main.go
*/

func main (){
	fmt.Println("showing math libary")
	mathlib.Sum(a,b)
	mathlib.Sub(a , b)
	mathlib.Mult(a , b)
	mathlib.Divi(a , b)

	fmt.Println("total age")
	
	totalage.Age(a , b)

}
