package models

// Estructura para representar un ítem
type Item struct {
	ID   int    `json:"id"`   // Identificador del ítem
	Name string `json:"name"` // Nombre del ítem

	// Nuevos campos
	CustomerName string `json:"customerName"` // Nombre del cliente que hizo el pedido
	OrderDate    string `json:"orderDate"`    // Fecha en que se hizo el pedido
	Product      string `json:"product"`      // Producto del pedido
	Quantity     int    `json:"quantity"`     // Cantidad del producto en el pedido
	Price        int    `json:"price"`        // Precio del producto en el pedido
}

// Lista de ítems, inicialmente vacía
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
