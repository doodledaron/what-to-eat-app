package handlers

import (
	"encoding/json"
	// "log"
	"net/http"
	"strconv"
	"what-to-eat-app/backend/models"
)

// in memory storage for items
var items []models.Food

// initialize the items
func Init() {
	items = []models.Food{
		{
			ID:   1,
			Name: "Pizza",
			Type: "Fast food",
		},
		{
			ID:   2,
			Name: "Burger",
			Type: "Fast food",
		},
		{
			ID:   3,
			Name: "Pasta",
			Type: "Main Course",
		},
	}
}

// get all items
func GetFood(w http.ResponseWriter, r *http.Request) {
	//set the header of the response
	w.Header().Set("Content-Type", "application/json")
	//encode the items to json and write it to the response
	json.NewEncoder(w).Encode(items)
}

// create an items
func CreateFood(w http.ResponseWriter, r *http.Request) {
	//set the header of the response
	w.Header().Set("Content-Type", "application/json")
	//create an empty item
	var item models.Food
	//decode the request body to the item
	json.NewDecoder(r.Body).Decode(&item)
	//append the item to the items
	items = append(items, item)
	//encode the item to json and write it to the response
	json.NewEncoder(w).Encode(item)
}

// delete an item
func DeleteFood(w http.ResponseWriter, r *http.Request) {
	//set the header of the response
	w.Header().Set("Content-Type", "application/json")
	//get the id from the request
	id := r.URL.Query().Get("id")
	//find the item with the id
	for i, item := range items {
		if strconv.Itoa(item.ID) == id {
			//delete the item from the items
			items = append(items[:i], items[i+1:]...)
			//encode the item to json and write it to the response
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	//if the item is not found, return a 404 status
	w.WriteHeader(http.StatusNotFound)
}

func UpdateFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Food
	// take the request body and decode it into the item which is a Food struct
	json.NewDecoder(r.Body).Decode(&item)

	//loop through the items and find the item with the same ID and update it
	//json needs to have the same ID and the updated fields
	for i, f := range items {
		if f.ID == item.ID {
			items[i] = item
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
