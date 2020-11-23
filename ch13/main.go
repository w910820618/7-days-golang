package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello,Go Web"))
}

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
