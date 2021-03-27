package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-api/models"
)

func HandleArticles(w http.ResponseWriter, req *http.Request) {
	RequestMethod := req.Method
	// logger
	log.Printf("%s\t%s\tHandleArticles", req.Method, req.RequestURI)
	if RequestMethod == "POST" {
		log.Println("POST")
	}
	if RequestMethod == "GET" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(models.Articles)
		if err != nil {
			fmt.Fprintf(w, "error to get the articles")
		}
	}
}
