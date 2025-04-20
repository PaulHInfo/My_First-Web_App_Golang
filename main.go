package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Users struct {
	Name  string `json:"name"` // nom en json
	Pswd  string `json:"-"`    // ignorer lors du JSON
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/test", test)
	http.HandleFunc("/search", search)
	http.HandleFunc("/login", login)
	http.HandleFunc("/user", user)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func user(w http.ResponseWriter, r *http.Request) {
	//fake data
	u := []Users{
		{Name: "bob", Pswd: "uriewu", Email: "test@gmail.com"},
		{Name: "alicia", Pswd: "kaka", Email: "etstasdf@gmail.com"},
	}
	//encodage

	b, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "index.html")
		return
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Print(err)
		}
		u := r.FormValue("Username")
		p := r.FormValue("password")
		fmt.Print(u)
		fmt.Print("   ")
		fmt.Print(p)
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
