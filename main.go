package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/database"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()
var db *sql.DB

type ArticlesFormData struct {
	Title, Body string
	URL         *url.URL
	Errors      map[string]string
}

type Article struct {
	Title, Body string
	ID          int64
}

type Object struct {
}

func (obj *Object) method() {

}

func function() {

}

func main() {
	database.Initialize()
	db = database.DB

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	// 通过命名路由获取 URL 示例
	homeURL, _ := router.Get("home").URL()
	fmt.Println("homeURL: ", homeURL)
	articleURL, _ := router.Get("articles.show").URL("id", "1")
	fmt.Println("articleURL: ", articleURL)

	http.ListenAndServe(":8000", middlewares.RemoveTrailingSlash(router))
}
