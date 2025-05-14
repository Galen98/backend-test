package main

import (
	"log"
	"logical-test/practical-test/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/summary", controller.GetSummary)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
