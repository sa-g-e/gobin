package main

import (
    "log"
    "net/http"
    "github.com/sa-g-e/gobin/handlers"
)

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", handlers.IndexHandler)
    http.HandleFunc("/create", handlers.CreateHandler)
    http.HandleFunc("/view/", handlers.ViewHandler)

    log.Println("server @ http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}