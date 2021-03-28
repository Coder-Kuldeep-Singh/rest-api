package main

import (
	"log"
	"net/http"
	"rest-api/routers"
)

func main() {
	router := routers.Router()
	log.Println(http.ListenAndServe(":8000", router))
}
