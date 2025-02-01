package main

import (
	"net/http"

	"project_sem/db"
	"project_sem/handlers"
)

func main() {

	connection, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	handler := &handlers.Handler{Connection: connection}

	mux := http.NewServeMux()
	mux.HandleFunc(`/api/v0/prices`, handler.PricesHandler)

	err = http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
