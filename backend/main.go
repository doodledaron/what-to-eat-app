package main

import (
	"log"
	"net/http"

	"what-to-eat-app/backend/handlers"
	"what-to-eat-app/backend/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

)

func main(){
	//initialize the router 
	router := mux.NewRouter()

	//initialize the routes
	//Init function from the handlers package to populate sample data
	handlers.Init()

	//initialize the routes
	routes.Init(router)

	//initialize the cors
	//CORS helps protect your web application by ensuring that only authorized domains can access your resources.
	//my nextjs can access the backend server which is local host8080 by using the cors
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	//wrap the handler with cors handler
	//enhance the router with additional middleware or functionality provided by CORS Handler
	//Middleware is a function that wraps an HTTP handler to add additional processing before or after the handler is executed.
	//For example, logging middleware might log each request, authentication middleware might check if the user is authenticated, and compression middleware might compress the response.
	handler := c.Handler(router)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
	//This function listens on the specified TCP network address and then calls the Serve method to handle incoming HTTP requests. If ListenAndServe encounters an error, it returns it, which is then passed to the Fatal method, causing the program to log the error and exit.
}