package main

import (
	"api"
	"net/http"
)

func main() {
	http.HandleFunc("/bmi", api.ApiHandler)
	http.ListenAndServe(":3000", nil)
}
