package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/test", test)
	http.HandleFunc("/search", search)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
	fmt.Fprint(w, "Adieu") // affiche sur la page
}
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
	fmt.Fprint(w, "test") // affiche sur la page
}

//requête http
/*
http://localhost:8000/search?t=go&p=1
t & p sont des paramètre
*/
func search(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("t")
	p := r.URL.Query().Get("p")
	fmt.Println("salut")
	fmt.Fprint(w, "salutttt")
	fmt.Fprint(w, t)
	fmt.Fprint(w, p)
}
