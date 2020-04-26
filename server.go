package main

import (
	"html/template"
	"log"
	"net/http"
)

func main(){
	log.Println("listen on 8080...")
	http.HandleFunc("/",index)
	//监听端口
	http.ListenAndServe(":8080", nil)
}
//输出首页
func index(w http.ResponseWriter, r *http.Request){
	t,_:=template.ParseFiles("./tmpl/index.html")
	t.Execute(w, nil)
}
