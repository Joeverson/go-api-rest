package main

import (
	"log"
	"net/http"

	"./modules/clients"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// fmt.Println(reflect.TypeOf(&router))
	clients.RegisterRoutes(router)

	log.Fatal(http.ListenAndServe(":8000", router))
}
