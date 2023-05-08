package controllers

import (
	"encoding/json"
	"fmt"
	"labora-api/models"
	"labora-api/services"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func Json(response http.ResponseWriter, status int, data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Errorf("error while mashalling object %v, trace: %+v", data, err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	_, err = response.Write(bytes)
	if err != nil {
		fmt.Errorf("error while writting bytes to response writer: %+v", err)
	}
}

func ObtenerItems(response http.ResponseWriter, _ *http.Request) {
	items, err := services.GetItems()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Error al obtener los items"))
		return
	}

	Json(response, http.StatusOK, items)
}

// Manejador de solicitudes HTTP GET para obtener la lista de elementos paginada
func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	// Establecer el tipo de contenido de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Obtener los parámetros de la consulta de URL
	params := r.URL.Query()
	page := params.Get("page")
	itemsPerPage := params.Get("itemsPerPage")

	// Asignar valores predeterminados si los parámetros no se proporcionan o son inválidos
	pageIndex, err := strconv.Atoi(page)
	if err != nil || pageIndex < 1 {
		pageIndex = 1
	}
	itemsPerPageInt, err := strconv.Atoi(itemsPerPage)
	if err != nil || itemsPerPageInt < 1 {
		itemsPerPageInt = 3
	}

	// Obtener la lista de elementos paginada
	newListItems, count, err := services.GetPaginatedItems(pageIndex, itemsPerPageInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Calcular el número total de páginas necesarias para mostrar todos los elementos
	totalPages := int(math.Ceil(float64(count) / float64(itemsPerPageInt)))

	// Crear un mapa que contiene información sobre la paginación
	paginationInfo := map[string]interface{}{
		"totalPages":  totalPages,
		"currentPage": pageIndex,
	}

	// Crear un mapa que contiene la lista de elementos y la información de paginación
	responseData := map[string]interface{}{
		"items":      newListItems,
		"pagination": paginationInfo,
	}

	// Codificar el mapa de respuesta en formato JSON y enviar en la respuesta HTTP
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func EditarItem(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idStr := vars["id"]
	// Convertir el ID de string a int

	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Manejar el error de la conversión
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("El ID debe ser un número"))
		return
	}
	//show id
	fmt.Println(id)

	var actualizacion models.Item
	//show request.body with struct names as object

	err = json.NewDecoder(request.Body).Decode(&actualizacion)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Error al procesar la solicitud"))
		return
	}

	for i, item := range models.Items {
		if item.ID == id {
			models.Items[i].Name = actualizacion.Name
			response.WriteHeader(http.StatusOK)
			json.NewEncoder(response).Encode(models.Items[i])
			return
		}
	}

	response.WriteHeader(http.StatusNotFound)
	response.Write([]byte("No se encontró el elemento con ID " + idStr))
}

func BuscarID(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	idStr := vars["id"]

	// Convertir el ID de string a int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Manejar el error de la conversión
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("El ID debe ser un número"))
		return
	}
	//id := getQueryParam(request, "id")

	// Buscar el elemento con el id que corresponde
	//var itemByID *Item
	var itemName string
	for i := 0; i < len(models.Items); i++ {
		if models.Items[i].ID == id {
			itemName = models.Items[i].Name
			break
		}
	}
	if itemName == "" {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("No se encontró el elemento con ID " + idStr))
		return

	}

	jsonData, err := json.Marshal(itemName)
	if err != nil {
		// Manejar el error de la conversión a JSON
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Error al convertir a JSON"))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte(jsonData))

}

func CrearItem(response http.ResponseWriter, request *http.Request) {
	var nuevoItem models.Item

	err := json.NewDecoder(request.Body).Decode(&nuevoItem)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Error al procesar la solicitud"))
		return
	}

	nuevoID := len(models.Items) + 1
	nuevoItem.ID = nuevoID
	models.Items = append(models.Items, nuevoItem)

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(nuevoItem)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var itemUpdate models.Item
	err := json.NewDecoder(r.Body).Decode(&itemUpdate)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	idDeItemComoTexto := vars["id"]
	idComoNumero, err := strconv.Atoi(idDeItemComoTexto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, item := range models.Items {
		if item.ID == idComoNumero {
			models.Items[i] = itemUpdate
			w.Write([]byte("Item actualizado correctamente"))
			return
		}
	}

	w.Write([]byte("No se pudo actualizar el item"))
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idDeItemComoTexto := vars["id"]
	idComoNumero, err := strconv.Atoi(idDeItemComoTexto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, item := range models.Items {
		if item.ID == idComoNumero {

			nuevoEstructuraDeItemsConUnItemMenos := append(models.Items[:i], models.Items[i+1:]...)

			models.Items = nuevoEstructuraDeItemsConUnItemMenos
			w.Write([]byte("Item eliminado correctamente"))
			return
		}
	}

	w.Write([]byte("Item no pudo ser eliminado"))
}

func GetItemByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	variables := r.URL.Query()
	name := variables.Get("name")

	for _, item := range models.Items {
		if strings.ToLower(item.Name) == strings.ToLower(name) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&models.Item{})
}
