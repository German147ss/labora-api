package services

import (
	"fmt"
	"labora-api/models"
)

func GetItems() ([]models.Item, error) {
	items := make([]models.Item, 0)
	rows, err := Db.Query("SELECT * FROM items")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

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
