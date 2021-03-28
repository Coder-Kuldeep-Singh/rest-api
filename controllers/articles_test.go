package controllers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"rest-api/models"
	"rest-api/routers"
	"strings"
	"testing"
)

func Client(method, url string, body io.Reader) (*httptest.ResponseRecorder, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	response := httptest.NewRecorder()
	routers.Router().ServeHTTP(response, request)
	return response, nil
}

func TestSearchTermWithValidArg(t *testing.T) {
	t.Log("..: INITIALIZED TestSearchTermWithValidArg case :..")
	url := "/articles/search?q=golang"
	method := http.MethodGet
	response, err := Client(method, url, nil)
	if err != nil {
		log.Printf("FAIL: TestSearchTermWithValidArg failed")
		t.Errorf("TestSearchTermWithValidArg failed to send the request %s\t%s", url, method)
	}
	if response.Code != http.StatusOK {
		log.Printf("FAIL: TestSearchTermWithValidArg failed")
		t.Errorf("Unexpected Status Code Found %d", response.Code)
	}
}

func TestSearchTermWithInValidArg(t *testing.T) {
	t.Log("..: INITIALIZED TestSearchTermWithInValidArg case :..")
	url := "/articles/search?q=testcases"
	method := http.MethodGet
	response, err := Client(method, url, nil)
	if err != nil {
		log.Printf("FAIL: TestSearchTermWithInValidArg failed")
		t.Errorf("TestSearchTermWithInValidArg failed to send the request %s\t%s", url, method)
	}
	if response.Code != http.StatusNotFound {
		log.Printf("FAIL: TestSearchTermWithInValidArg failed")
		t.Errorf("Unexpected Status Code Found %d", response.Code)
	}
}

func TestSearchTermWithBlankArg(t *testing.T) {
	t.Log("..: INITIALIZED TestSearchTermWithBlankArg case :..")
	url := "/articles/search?q="
	method := http.MethodGet
	response, err := Client(method, url, nil)
	if err != nil {
		log.Printf("FAIL: TestSearchTermWithBlankArg failed")
		t.Errorf("TestSearchTermWithBlankArg failed to send the request %s\t%s", url, method)
	}
	if response.Code != http.StatusNotAcceptable {
		log.Printf("FAIL: TestSearchTermWithBlankArg failed")
		t.Errorf("Unexpected Status Code Found %d", response.Code)
	}
}

func TestGetArticleByIDWithValidArg(t *testing.T) {
	t.Log("..: INITIALIZED TestGetArticleByIDWithValidArg case :..")
	url := "/articles/1"
	method := http.MethodGet
	response, err := Client(method, url, nil)
	if err != nil {
		log.Printf("FAIL: TestGetArticleByIDWithValidArg failed")
		t.Errorf("TestGetArticleByIDWithValidArg failed to send the request %s\t%s", url, method)
	}
	if response.Code != http.StatusOK {
		log.Printf("FAIL: TestGetArticleByIDWithValidArg failed")
		t.Errorf("Unexpected Status Code Found %d", response.Code)
	}
}

func TestGetArticleByIDInValidArg(t *testing.T) {
	t.Log("..: INITIALIZED TestGetArticleByIDInValidArg case :..")
	url := "/articles/3"
	method := http.MethodGet
	response, err := Client(method, url, nil)
	if err != nil {
		log.Printf("FAIL: TestGetArticleByIDInValidArg failed")
		t.Errorf("TestGetArticleByIDInValidArg failed to send the request %s\t%s", url, method)
	}
	if response.Code != http.StatusNotFound {
		log.Printf("FAIL: TestGetArticleByIDInValidArg failed")
		t.Errorf("Unexpected Status Code Found %d", response.Code)
	}
}

// func TestGetArticleByIDWithStringArg(t *testing.T) {
// 	t.Log("..: INITIALIZED TestGetArticleByIDWithStringArg case :..")
// 	url := "/articles/`1`"
// 	method := http.MethodGet
// 	response, err := Client(method, url, nil)
// 	if err != nil {
// 		log.Printf("FAIL: TestGetArticleByIDWithStringArg failed")
// 		t.Errorf("TestGetArticleByIDWithStringArg failed to send the request %s\t%s", url, method)
// 	}
// 	if response.Code != http.StatusNotAcceptable {
// 		log.Printf("FAIL: TestGetArticleByIDWithStringArg failed")
// 		t.Errorf("Unexpected Status Code Found %d", response.Code)
// 	}
// }

func TestPostArticleWithoutData(t *testing.T) {
	t.Log("..: INITIALIZED TestPostArticleWithNoData case :..")
	url := "/articles"
	method := http.MethodPost
	response, err := Client(method, url, nil)
	if err != nil {
		log.Printf("FAIL: TestPostArticleWithNoData failed")
		t.Errorf("TestPostArticleWithNoData failed to send the request %s\t%s", url, method)
	}
	if response.Code != http.StatusBadRequest {
		log.Printf("FAIL: TestPostArticleWithNoData failed")
		t.Errorf("Unexpected Status Code Found %d", response.Code)
	}
}

func TestPostArticleWithObject(t *testing.T) {
	t.Log("..: INITIALIZED TestPostArticleWithObject case :..")
	url := "/articles"
	method := http.MethodPost
	article := models.Article{
		Id:       2,
		Title:    "Learning Golang",
		Content:  "Go",
		SubTitle: "Golang",
	}
	MarshalledArticle, _ := json.Marshal(&article)
	response, err := Client(method, url, bytes.NewBuffer(MarshalledArticle))
	if err != nil {
		log.Printf("FAIL: TestPostArticleWithObject failed")
		t.Errorf("TestPostArticleWithObject failed to send the request %s\t%s", url, method)
	}
	want := "id already exists"
	if !strings.Contains(response.Body.String(), want) {
		log.Printf("FAIL: TestPostArticleWithObject failed")
		t.Errorf("Unexpected Status Code Found %s", response.Body.String())
	}

}

func TestPostArticleWithArrayObjects(t *testing.T) {
	t.Log("..: INITIALIZED TestPostArticleWithArrayObjects case :..")
	url := "/articles"
	method := http.MethodPost
	article := []models.Article{
		{
			Id:       2,
			Title:    "Learning Golang",
			Content:  "Go",
			SubTitle: "Golang",
		},
		{
			Id:       3,
			Title:    "Learning Golang",
			Content:  "Go",
			SubTitle: "Golang",
		},
	}
	MarshalledArticle, _ := json.Marshal(&article)
	response, err := Client(method, url, bytes.NewBuffer(MarshalledArticle))
	if err != nil {
		t.Errorf("TestPostArticleWithArrayObjects failed to send the request %s\t%s", url, method)
	}
	got := strings.Contains(response.Body.String(), "id already exists")
	if !got {
		t.Errorf(response.Body.String())
		t.Errorf("Unexpected Status Code Found %s", response.Body.String())
	}
}

func TestPostArticle(t *testing.T) {
	t.Log("..: INITIALIZED TestPostArticle case :..")
	url := "/articles"
	method := http.MethodPost
	article := models.Article{
		Title:    "Learning Golang",
		Content:  "Go",
		SubTitle: "Golang",
	}
	MarshalledArticle, _ := json.Marshal(&article)
	response, err := Client(method, url, bytes.NewBuffer(MarshalledArticle))
	if err != nil {
		log.Printf("FAIL:TestPostArticle failed")
		t.Errorf("TestPostArticle failed to send the request %s\t%s", url, method)
	}
	want := "request processed successfully"
	if !strings.Contains(response.Body.String(), want) {
		log.Printf("FAIL: TestPostArticle failed")
		t.Errorf("Unexpected Status Code Found %s", response.Body.String())
	}
}

func TestGetArticles(t *testing.T) {
	t.Log("..: INITIALIZED TestGetArticals case :..")
	url := "/articles"
	method := http.MethodGet
	response, err := Client(method, url, nil)
	if err != nil {
		log.Printf("FAIL: TestGetArticals failed")
		t.Errorf("TestGetArticals failed to send the request %s\t%s", url, method)
	}
	if response.Code != http.StatusOK {
		log.Printf("FAIL: TestGetArticals failed")
		t.Errorf("Unexpected Status Code Found %d", response.Code)
	}
}

func TestGetArticalWithArgs(t *testing.T) {
	t.Log("..: INITIALIZED TestGetArticalWithArgs case :..")
	url := "/articles?limit=2&cursor=1"
	method := http.MethodGet
	response, err := Client(method, url, nil)
	if err != nil {
		log.Printf("FAIL: TestGetArticalWithArgs failed")
		t.Errorf("TestGetArticalWithArgs failed to send the request %s\t%s", url, method)
	}
	if response.Code != http.StatusOK {
		log.Printf("FAIL: TestGetArticalWithArgs failed")
		t.Errorf("Unexpected Status Code Found %d", response.Code)
	}
}
