package main

import "net/http"

func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":8888", nil)
}

func HelloWorld(w http.ResponseWritter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}