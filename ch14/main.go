package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func GetForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("/Users/wuzhouyang/go/src/7-days-golang/ch14/login.html")
		t.Execute(w, nil)
	} else {
		fmt.Printf("username: %s\n", r.Form["username"][0])
		fmt.Printf("password: %s\n", r.Form["password"][0])
		if r.Form["username"][0] != "root" || r.Form["password"][0] != "123" {
			t, _ := template.ParseFiles("/Users/wuzhouyang/go/src/7-days-golang/ch14/error.html")
			fmt.Println("登陆失败")
			t.Execute(w, nil)
		} else {
			t, _ := template.ParseFiles("/Users/wuzhouyang/go/src/7-days-golang/ch14/success.html")
			fmt.Println("登陆成功")
			t.Execute(w, nil)
		}
	}
}

func main() {
	http.HandleFunc("/", GetForm)

	err := http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
