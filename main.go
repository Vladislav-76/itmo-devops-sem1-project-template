package main

import (
    "net/http"

    "project_sem/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/api/v0/prices`, handlers.PricesHandler)

    err := http.ListenAndServe(`:8080`, mux)
    if err != nil {
        panic(err)
    }
}
