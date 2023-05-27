package main

import (
	"log"
	"net/http"
	"text/template"

	"fastIM/app/controller"
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
	http.HandleFunc("/user/login", controller.UserLogin)
	http.HandleFunc("/user/register", controller.UserRegister)
	http.HandleFunc("/contact/add_friend", controller.AddFriend)
	http.HandleFunc("/contact/loadfriend", controller.LoadFriend)
	http.HandleFunc("/contact/createcommunity", controller.CreateCommunity)
	http.HandleFunc("/contact/joincommunity", controller.JoinCommunity)
	http.HandleFunc("/chat", controller.Chat)

	// 提过静态资源目录支持
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	http.Handle("/resource", http.FileServer(http.Dir(".")))
	registerView()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
