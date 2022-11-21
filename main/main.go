package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nawazish-github/vendor/controllers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/vendor", controllers.VendorHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}