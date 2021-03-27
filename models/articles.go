package models

import "time"

var CurrentID int

type Article struct {
	Id        int
	Title     string
	Content   string
	SubTitle  string
	CreatedAt time.Time
}

// make Article accessable for all the packages for additional use(like get article by id,get all articles from memory)
var Articles []Article

// Give us some seed data
func init() {
	CreateArticle(Article{Title: "Learning Golang", Content: "Go", SubTitle: "Golang"})
	CreateArticle(Article{Title: "Golang Data Types", Content: "Go", SubTitle: "Golang"})
}

func GetArticleByID(id int) Article {
	for _, article := range Articles {
		if article.Id == id {
			return article
		}
	}
	// return empty Todo if not found
	return Article{}
}

func CreateArticle(article Article) Article {
	CurrentID += 1
	article.Id = CurrentID
	Articles = append(Articles, article)
	return article
}
