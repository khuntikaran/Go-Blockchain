package main

import (
	"blockchain/chainS"
	"blockchain/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	chainS.Init()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/blockchain", handlers.GetBlockchain).Methods("GET")
	router.HandleFunc("/add", handlers.CreateBlock).Methods("POST")
	log.Fatal(http.ListenAndServe(":5050", router))
}
