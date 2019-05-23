package webServer

import (
	"html/template"
	"net/http"
)

// 控制器
func handler(w http.ResponseWriter, r *http.Request) {
	tem, _ := template.ParseFiles("../demo/webServer/temp.html")
	tem.Execute(w, r.URL.Path)
	// fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
