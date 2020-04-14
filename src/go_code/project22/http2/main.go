package main

import (
    "fmt"
    "net/http"
)

type indexHandler struct {
    content string
}

func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, ih.content)
}

func main() {
	/*  http.HandleFunc 和 http.Handle都是用于注册路由，可以发现两者的区别
		在于第二个参数前者是一个具有
			func(w http.ResponseWriter, r *http.Requests)
		签名的"函数"，而后者是一个"结构体"，该结构体实现了
			func(w http.ResponseWriter, r *http.Requests)
		签名的方法。
		
	*/
	http.Handle("/", &indexHandler{content: "hello world!"})
	

    http.ListenAndServe(":8001", nil)
}