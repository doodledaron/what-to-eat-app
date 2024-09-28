package routes

import (
	"what-to-eat-app/backend/handlers"
	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	router.HandleFunc("/api/v1/get-food", handlers.GetFood).Methods("GET")

	router.HandleFunc("/api/v1/create-food", handlers.CreateFood).Methods("POST")

	router.HandleFunc("/api/v1/delete-food", handlers.DeleteFood).Methods("DELETE")

	router.HandleFunc("/api/v1/update-food", handlers.UpdateFood).Methods("PUT")
}