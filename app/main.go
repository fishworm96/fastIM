package main

import (
	"log"
	"net/http"
	"text/template"
)

func registerView() {
	tpl, err := template.ParseGlob("./app/view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplName := v.Name()
		http.HandleFunc(tplName, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tplName, nil)
		})
	}
}

func main() {
	// 提过静态资源目录支持
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	http.Handle("/resource", http.FileServer(http.Dir(".")))
	registerView()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
