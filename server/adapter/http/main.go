package main

import (
	"net/http"

	"github.com/rafaelsouzaribeiro/client-serve-api-go/server/adapter/sqllite"

	"github.com/rafaelsouzaribeiro/client-serve-api-go/server/adapter/http/di"

	"github.com/gorilla/mux"
)

func main() {
	conn := sqllite.Connect()
	defer conn.Close()

	priceService := di.ConfigProductDI(conn)

	router := mux.NewRouter()
	router.HandleFunc("/cotacao", priceService.Get).Methods("GET")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		panic(err)
	}

}
