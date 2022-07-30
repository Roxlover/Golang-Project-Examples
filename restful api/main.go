package main

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page"))

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

func main() {
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8080", nil)
}
