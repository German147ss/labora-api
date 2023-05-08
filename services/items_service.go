package services

import (
	"fmt"
	"labora-api/models"
)

// GetItems obtiene todos los items de la tabla 'items' de la base de datos.
// Retorna una lista de struct 'models.Item' y un error en caso de que haya ocurrido alguno.
func GetItems() ([]models.Item, error) {
	items := make([]models.Item, 0)
	rows, err := Db.Query("SELECT * FROM items")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	// Itera sobre cada fila en 'rows' y crea una instancia de 'models.Item' con los valores de cada columna.
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		items = append(items, item)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return items, nil
}

func GetPaginatedItems(pageIndex, itemsPerPage int) ([]models.Item, int, error) {
	//log params
	fmt.Printf("pageIndex: %d, itemsPerPage: %d\n", pageIndex, itemsPerPage)
	// Calcular el índice inicial y el límite de elementos en función de la página actual y los elementos por página
	startIndex := (pageIndex - 1) * itemsPerPage

	// Obtener el número total de filas en la tabla items
	var count int
	err := Db.QueryRow("SELECT COUNT(*) FROM items").Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	// Obtener la lista de elementos correspondientes a la página actual
	rows, err := Db.Query("SELECT * FROM items ORDER BY id OFFSET $1 LIMIT $2", startIndex, itemsPerPage)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var newListItems []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price)
		if err != nil {
			return nil, 0, err
		}
		newListItems = append(newListItems, item)
	}

	if len(newListItems) == 0 {
		return nil, 0, fmt.Errorf("No items found for page %d", pageIndex)
	}

	return newListItems, count, nil
}

// UpdateItemByID actualiza un item en la base de datos a partir de su ID
func UpdateItemByID(id int, item models.Item) (models.Item, error) {
	var updatedItem models.Item
	row := Db.QueryRow("UPDATE items SET customer_name = $1, order_date = $2, product = $3, quantity = $4, price = $5 WHERE id = $6 RETURNING *",
		item.CustomerName, item.OrderDate, item.Product, item.Quantity, item.Price, id)
	err := row.Scan(&updatedItem.ID, &updatedItem.CustomerName, &updatedItem.OrderDate, &updatedItem.Product, &updatedItem.Quantity, &updatedItem.Price)
	if err != nil {
		fmt.Println(err)
		return updatedItem, err
	}
	return updatedItem, nil
}
