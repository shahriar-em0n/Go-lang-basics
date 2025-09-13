package main

type Cars struct {
	Id       int     `json:"id"`
	Brand    string  `json:"brand"`
	Model    string  `json:"model"`
	Color    string  `json:"color"`
	Price    float32 `json:"price"`
	ImageUrl string  `json:"imageUrl"`
}