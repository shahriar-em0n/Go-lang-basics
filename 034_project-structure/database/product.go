package database

type Car struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`

	Brand       string `json:"brand"`
	Color       string `json:"color"`
	Year        int    `json:"year"`
	IsAvailable bool   `json:"isAvailable"`
}

var CarList []Car

