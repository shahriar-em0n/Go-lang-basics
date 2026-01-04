package main

import (
	"ecom/cmd"
	"ecom/database"
)

func main() {
	cmd.Server()
}

func init() {
	cars := []database.Car{
		{
			ID:          1,
			Title:       "2020 Toyota Camry",
			Description: "Reliable sedan with excellent fuel economy. Perfect for daily commuting.",
			Price:       25000.00,
			ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRVL4LrER0QF9hJvJrh_6Z91uPCHhyoiSH93Q&s",
			Brand:       "Toyota",
			Year:        2020,
			Color:       "Silver",
			IsAvailable: true,
		},
		{
			ID:          2,
			Title:       "2019 Honda Civic",
			Description: "Sporty compact car with modern features. Low mileage, well maintained.",
			Price:       22000.00,
			ImgUrl:      "https://www.zimbrickfishhatcheryroad.com/blogs/2133/wp-content/uploads/2019/06/2019-honda-civic-madison-wi-1.jpg",
			Brand:       "Honda",
			Year:        2019,
			Color:       "Red",
			IsAvailable: true,
		},
		{
			ID:          3,
			Title:       "2021 Ford Mustang",
			Description: "Powerful sports car with V8 engine. Thrilling driving experience.",
			Price:       45000.00,
			ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTHmo7dbLpKH2TjrWiaCdcWnlkibgBPT55ANA&s",
			Brand:       "Ford",
			Year:        2021,
			Color:       "Red",
			IsAvailable: true,
		},
		{
			ID:          4,
			Title:       "2022 Tesla Model 3",
			Description: "Electric sedan with autopilot. Zero emissions, amazing technology.",
			Price:       55000.00,
			ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSW7aMb119tpQN6MFrNgr3s-QuXgfyzdH3_xQ&s",
			Brand:       "Tesla",
			Year:        2022,
			Color:       "White",
			IsAvailable: false,
		},
		{
			ID:          5,
			Title:       "2018 BMW X5",
			Description: "Luxury SUV with premium interior. Spacious and comfortable.",
			Price:       38000.00,
			ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcThgiE6XljTYo6op5IhfJZxLhPQv7zdgbQhjA&s",
			Brand:       "BMW",
			Year:        2018,
			Color:       "Black",
			IsAvailable: true,
		},
		{
			ID:          6,
			Title:       "2020 Mazda CX-5",
			Description: "Compact SUV with great handling. Ideal for small families.",
			Price:       28000.00,
			ImgUrl:      "https://media.drive.com.au/obj/tx_rs:auto:4928:3334:1/driveau/upload/cms/uploads/pyryCDwzQBEwiUQEQePa",
			Brand:       "Mazda",
			Year:        2020,
			Color:       "Gray",
			IsAvailable: true,
		},
	}

	database.CarList = append(database.CarList, cars...)
}
