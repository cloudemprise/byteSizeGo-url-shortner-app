package main

import (
	"log"
	"net/http"

	"github.com/cloudemprise/byteSizeGo-url-shortner-app/internal/controllers"
)

func main() {
	http.HandleFunc("/", controllers.ShowIndex)
	http.HandleFunc("/shorten", controllers.Shorten)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
