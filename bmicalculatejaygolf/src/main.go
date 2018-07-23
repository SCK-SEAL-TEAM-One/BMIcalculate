package main

import (
	"api"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/bmi/", http.StripPrefix("/bmi/", fileServer))
	http.HandleFunc("/bmi/calculate", api.BmiHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
