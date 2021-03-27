package models_test

import (
	"log"
	"rest-api/models"
	"testing"
)

func TestCreateArticleByEmptyArg(t *testing.T) {
	t.Log("..: INITIALIZED TestCreateArticleByEmptyArg case :..")
	article := models.Article{}
	generated, err := models.CreateArticle(article)
	if err != nil {
		t.Errorf("FAIL: error to generate article: %s", err.Error())
	}
	// log.Println(generated)
	if generated.Id != 3 {
		t.Error("FAIL: TestCreateArticleByEmptyArg test case failed to generate the article")
	}
}

func TestCreateArticleByValidArgs(t *testing.T) {
	t.Log("..: INITIALIZED TestCreateArticleByValidArgs case :..")
	article := models.Article{Title: "Learning Golang", Content: "Go", SubTitle: "Golang"}
	generated, err := models.CreateArticle(article)
	if err != nil {
		t.Errorf("FAIL: error to generate article: %s", err.Error())
	}
	if generated.Id != 3 {
		t.Error("FAIL: TestCreateArticleByValidArgs test case failed to generate the article")
	}
}

func TestNotExistsIDArg(t *testing.T) {
	t.Log("..: INITIALIZED TestGetArticleByID case :..")
	want := 0
	article := models.GetArticleByID(want)
	if article.Id != want {
		t.Error("FAIL: error to get article by id")
	} else {
		log.Println("PASS: test case pass")
	}
}

func TestExistsArticleIDArg(t *testing.T) {
	t.Log("..: INITIALIZED TestExistsArticleIDArg case :..")
	want := 2
	article := models.GetArticleByID(want)
	if article.Id != want {
		t.Errorf("FAIL: error to get article by id: %d", want)
	} else {
		log.Println("PASS: test case pass")
	}
}

func TestSearchArticleByEmptyArg(t *testing.T) {
	t.Log("..: INITIALIZED TestSearchArticleByEmptyArg case :..")
	articles := models.SearchArticles("")
	if len(articles) != 0 {
		t.Error("FAIL: unexpected data found, Expecting nil articles")
	} else {
		log.Println("PASS: test case pass")
	}
}

func TestSearchArticleByArg(t *testing.T) {
	t.Log("..: INITIALIZED TestSearchArticleByArg case :..")
	articles := models.SearchArticles("data types")
	if len(articles) == 0 {
		t.Error("FAIL: unexpected data found")
	}

	if articles[0].Id != 2 {
		t.Errorf("FAIL: unexpected result found %+v", articles)
	} else {
		log.Println("PASS: test case pass")
	}
}
