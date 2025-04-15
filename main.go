package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
	fmt.Fprint(w, "Adieu") // affiche sur la page

}
