package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description" `
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

var (
	productList []Product
)

func getProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet { //  r.method = post, put, patch, delete
		http.Error(w, "Please give me GET request", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/products", getProducts)

	fmt.Println("Server running on :3000")

	err := http.ListenAndServe(":3000", mux) // error value nil

	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red. I love orange",
		Price:       100,
		ImageUrl:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT-6GoLpPXoOkT0lAuFcxXwJSQ7nxtRQqVJLg&s",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is red. I love Apple",
		Price:       100,
		ImageUrl:    "https://media.istockphoto.com/id/185262648/photo/red-apple-with-leaf-isolated-on-white-background.jpg?s=612x612&w=0&k=20&c=gUTvQuVPUxUYX1CEj-N3lW5eRFLlkGrU_cwwwOWxOh8=",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is Yellow. I love Banana",
		Price:       100,
		ImageUrl:    "https://img.freepik.com/free-vector/simple-isolated-banana_1308-125007.jpg?semt=ais_incoming&w=740&q=80",
	}

	prd4 := Product{
		ID:          4,
		Title:       "Mango",
		Description: "Mango is Yellow. I love Mango",
		Price:       100,
		ImageUrl:    "https://img.freepik.com/free-photo/mango-still-life_23-2151542114.jpg?w=360",
	}

	prd5 := Product{
		ID:          5,
		Title:       "lichi",
		Description: "lichi is Yellow. I love lichi",
		Price:       100,
		ImageUrl:    "https://m.media-amazon.com/images/I/81F6bQaa2NL._UF1000,1000_QL80_.jpg",
	}

	prd6 := Product{
		ID:          6,
		Title:       "Banana",
		Description: "Banana is Yellow. I love Banana",
		Price:       100,
		ImageUrl:    "https://img.freepik.com/free-vector/simple-isolated-banana_1308-125007.jpg?semt=ais_incoming&w=740&q=80",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
	productList = append(productList, prd6)

}
