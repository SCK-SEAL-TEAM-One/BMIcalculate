package main

import (
    "net/http"
    "log"
)


func main() {
    fs := http.FileServer(http.Dir("./public"))
    http.Handle("/bmi/", http.StripPrefix("/bmi/", fs))
    log.Fatal(http.ListenAndServe(":4321", nil))
}