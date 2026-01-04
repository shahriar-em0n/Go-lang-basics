package cmd

import (
	grobalrouter "ecom/grobal_router"
	"ecom/handlers"
	"ecom/middleware"
	"fmt"
	"net/http"
)

func Server() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.PreFlight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.wrappedMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error starting on the server", err)
	}
}
