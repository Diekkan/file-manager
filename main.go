package main

import (
	"filemanager/router"
	"log"
	"net/http"
)

func main() {
	r := router.SetupRouter()
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
