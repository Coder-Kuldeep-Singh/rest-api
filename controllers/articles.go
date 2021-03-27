package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"rest-api/models"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// GetArticles renders the articles
func GetArticles(w http.ResponseWriter, req *http.Request) {
	// logger
	log.Printf("%s\t%s\tGetArticles", req.Method, req.RequestURI)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(models.Articles)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprintf(w, "error to get the articles")
	}
}

func PostArticles(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s\t%s\tPostArticles", req.Method, req.RequestURI)
	if req.Body == nil {
		log.Printf("request body is missing {%d}\n", http.StatusBadRequest)
		fmt.Fprintf(w, "response body is nil")
		return
	}

	if !strings.Contains(req.Header.Get("Content-type"), "application/json") {
		log.Printf("content type not accepted {%s}", req.Header.Get("Content-type"))
		fmt.Fprintf(w, "not valid content-type")
		return
	}

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("error to read post data{Error:%v}", err)
	}
	defer req.Body.Close()

	x := bytes.TrimLeft(data, " \t\r\n")
	isArray := len(x) > 0 && x[0] == '['
	isObject := len(x) > 0 && x[0] == 123

	// r.Body is a io.ReadCloser, ioutil.NopCloser() in conjunction with bytes.NewReader()
	// Note :- https://stackoverflow.com/questions/53359013/read-request-body-two-times-in-golang
	reader := ioutil.NopCloser(bytes.NewReader(data))

	if isArray {
		HandlePostResponse(ResponseIsArray, reader, "ResponseIsArray", w)
		return
	}
	if isObject {
		HandlePostResponse(ResponseIsObject, reader, "ResponseIsObject", w)
		return
	}
}

func HandlePostResponse(PostedData func(buf io.Reader) error, response io.ReadCloser, funcName string, w http.ResponseWriter) {
	err := PostedData(response)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("%s", err.Error()))
		return
	}
	fmt.Fprintf(w, "request processed successfully")
}

// ResponseIsArray returns []LeadsInfo,err
func ResponseIsArray(buf io.Reader) error {
	var articles []models.Article
	err := json.NewDecoder(buf).Decode(&articles)
	if err != nil {
		return fmt.Errorf("json decoding failed isArray {stautsCode:%d} {Error:%v}", http.StatusBadRequest, err)
	}
	var Failed string
	for _, article := range articles {
		_, err = models.CreateArticle(article)
		if err != nil {
			Failed += fmt.Sprintf("%s\n", err.Error())
			log.Printf("ResponseIsArray failed:  %s", err.Error())
			continue
		}
	}
	return fmt.Errorf("%s", Failed)
}

// ResponseIsObject returns LeadsInfo,err
func ResponseIsObject(buf io.Reader) error {
	var article models.Article
	err := json.NewDecoder(buf).Decode(&article)
	if err != nil {
		return fmt.Errorf("json decoding failed isArray {stautsCode:%d} {Error:%v}", http.StatusBadRequest, err)
	}
	_, err = models.CreateArticle(article)
	if err != nil {
		return fmt.Errorf("ResponseIsObject failed: %s", err.Error())
	}
	return nil
}

// GetArticlesByID  returns  the matched id article
func GetArticleByID(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s\t%s\tGetArticleByID", req.Method, req.RequestURI)
	ID := mux.Vars(req)["id"]
	articleID, err := strconv.Atoi(ID)
	if err != nil {
		log.Printf("GetArticles : error to type cast the article ID: %s", err.Error())
		fmt.Fprintf(w, "request failed")
	}
	article := models.GetArticleByID(articleID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "error to get the articles")
	}
}

// SearchTerm renders the article on the search based
func SearchTerm(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s\t%s\tSearchTerm", req.Method, req.RequestURI)
	keys, ok := req.URL.Query()["q"]
	searchedQuery := ""
	if ok {
		searchedQuery = keys[0]
	}
	if searchedQuery == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "query is blank")
		return
	}
	articles := models.SearchArticles(strings.ToLower(searchedQuery))
	if len(articles) == 0 {
		fmt.Fprintf(w, "No records found")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(w).Encode(articles)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "error to get the articles")
		return
	}
}
