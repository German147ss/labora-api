package models

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Items []Item = []Item{
	{
		ID:   1,
		Name: "Camila",
	},
	{
		ID:   2,
		Name: "Paula",
	},
	{
		ID:   3,
		Name: "Alejandra",
	},
	{
		ID:   4,
		Name: "Andres",
	},
	{
		ID:   5,
		Name: "Luis",
	},
	{
		ID:   6,
		Name: "Camilo",
	},
	{
		ID:   7,
		Name: "Luisa",
	},
	{
		ID:   8,
		Name: "Juan",
	},
	{
		ID:   9,
		Name: "Liz",
	},
	{
		ID:   10,
		Name: "Carmen",
	},
}