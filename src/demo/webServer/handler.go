package webServer

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
)

// 控制器
func handler(w http.ResponseWriter, r *http.Request) {
	tem, _ := template.ParseFiles("../demo/webServer/temp/index.html")
	tem.Execute(w, r.URL.Path)
	// fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func loginHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		tem, _ := template.ParseFiles("../demo/webServer/temp/form.html")
		tem.Execute(w, nil)
	} else {
		fmt.Println(r.URL)
		r.ParseForm()
		fmt.Println("username:", r.Form["username"][0])
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("password:", r.Form["password"][0])

		if len(r.Form.Get("username")) == 0 {
			fmt.Fprintln(w, "no username")
		} else {
			fmt.Fprintln(w, "has")
			getint, err := strconv.Atoi(r.Form.Get("username"))
			if err != nil {
				fmt.Fprintln(w, "no num")
			}

			if getint > 100 {
				fmt.Fprintln(w, "bigger than 100")
			}
		}
	}
}
func uploadHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("../demo/webServer/temp/upload.html")
		t.Execute(w, nil)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("../demo/test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
