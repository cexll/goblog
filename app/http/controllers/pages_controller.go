package controllers

import (
	"fmt"
	"goblog/pkg/view"
	"net/http"
)

// PagesController 处理静态页面
type PagesController struct {
}

func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog! </h1>")
}

func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	view.Render(w, nil, "pages.about")
	// fmt.Fprint(w, "此博客是用以记录编程笔记,如你有反馈或建议,请联系 "+
	// 	"<a href=\"mailto:evanxian9@gmail.com\">evanxian9@gmail.com</a>")
}

func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// w.WriteHeader(http.StatusNotFound)
	// fmt.Fprint(w, "<h1>请求页面未找到 :( </h1>"+
	// 	"<p>如有疑惑, 请联系我们. </p>")
	view.Render(w, nil, "pages.404")
}
