package main

import (
    "fmt"
    "net/http"
)


func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello world")
}

func main() {
	/*利用http.HandleFunc在根路由/上注册了一个indexHandler*/
	http.HandleFunc("/", indexHandler)
	// 利用http.ListenAndServe开启监听
	http.ListenAndServe(":8000", nil)
}