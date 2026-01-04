package grobalrouter

import "net/http"


func GlobalRouter(mux http.Handler) http.Handler {
	handlAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-type")
		w.Header().Set("Access-Control-Allow-Method", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}

		mux.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handlAllReq)
}