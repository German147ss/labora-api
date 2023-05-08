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

// Obtener la lista de elementos paginada
func GetPaginatedItems(pageIndex, itemsPerPage int) ([]models.Item, error) {
	// Calcular el índice inicial y el límite de elementos en función de la página actual y los elementos por página
	startIndex := (pageIndex - 1) * itemsPerPage
	endIndex := startIndex + itemsPerPage

	// Obtener la lista de elementos correspondientes a la página actual
	var newListItems []models.Item
	if startIndex < len(models.Items) {
		if endIndex > len(models.Items) {
			newListItems = models.Items[startIndex:]
		} else {
			newListItems = models.Items[startIndex:endIndex]
		}
	}

	if len(newListItems) == 0 {
		return nil, fmt.Errorf("No items found for page %d", pageIndex)
	}

	return newListItems, nil
}
