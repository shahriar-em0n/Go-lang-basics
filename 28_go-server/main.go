package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello world")
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "I'm Shahriar. I'm Software engineer")
}

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)

	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("Server running on :3000")

	err := http.ListenAndServe(":3000", mux) // error value nil

	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}
