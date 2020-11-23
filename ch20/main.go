package main

import (
	"fmt"
	"net/http"
)

func HandlerToken(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Header.Get("AuthToken"))
}

func main() {
	http.HandleFunc("/token", HandlerToken)

	http.ListenAndServe(":8080", nil)
}
