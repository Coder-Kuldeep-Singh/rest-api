package routers

import (
	"rest-api/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	// r.HandleFunc("/articles", controllers.GetArticles).Methods("GET")
	r.Path("/articles").HandlerFunc(controllers.GetArticles).Methods("GET").Name("GetArticles")
	// r.HandleFunc("/articles", controllers.PostArticles).Methods("POST")
	r.Path("/articles").HandlerFunc(controllers.PostArticles).Methods("POST").Name("PostArticles")
	// r.HandleFunc("/articles/{id:[0-9]+}", controllers.GetArticleByID).Methods("GET")
	r.Path("/articles/{id:[0-9]+}").HandlerFunc(controllers.GetArticleByID).Methods("GET").Name("GetArticles")
	r.Path("/articles/search").HandlerFunc(controllers.SearchTerm).Methods("GET").Name("SearchTerm")
	return r
}
