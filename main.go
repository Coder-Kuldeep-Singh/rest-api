package main

import (
	"log"
	"net/http"
	"rest-api/controllers"
)

func main() {
	http.HandleFunc("/articles", controllers.HandleArticles)
	log.Println(http.ListenAndServe(":8000", nil))
}
