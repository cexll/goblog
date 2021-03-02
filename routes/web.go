package routes

import (
	"net/http"

	"goblog/app/http/controllers"
	"github.com/gorilla/mux"
)

func RegisterWebRoutes(r *mux.Router)  {
	// 静态页面
	pc := new(controllers.PagesController)
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")
}