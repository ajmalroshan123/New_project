package main

import (
	httphandlers "datapipeline/internal/http-handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Endpoints
	router.HandleFunc("/", httphandlers.RootHandler).Methods("GET")
	router.HandleFunc("/api/user/signup", httphandlers.Signup).Methods("POST")
	router.HandleFunc("/api/user/login", httphandlers.Login).Methods("POST")

	router.HandleFunc("/api/favourites/add", httphandlers.AddToFavouritesHandler).Methods("POST")
	router.HandleFunc("/api/favourites/remove", httphandlers.RemoveFromFavouritesHandler).Methods("DELETE")
	router.HandleFunc("/api/favourites/get", httphandlers.GetFavouritesHandler).Methods("GET")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
