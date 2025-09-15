package cmd

import (
	"eccomarce/global_router"
	"eccomarce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	// mux.Handle("GET /wlc", http.HandlerFunc(welcome))

	mux.Handle("GET /products",  (http.HandlerFunc(handlers.Getcars)))

	mux.Handle("POST /products",  (http.HandlerFunc(handlers.AddCars)))

	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductById))

	fmt.Println("Server running on :8080")

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Server error starting", err)
	}
}

