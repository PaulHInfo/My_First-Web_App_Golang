package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var (
	counter int
	mu      sync.Mutex // -
)

func main() {
	http.HandleFunc("/", servePage)
	http.HandleFunc("/increment", incrementCounter)
	http.HandleFunc("/parle", parle)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func servePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counter++
	if counter > 12 {
		counter = 0
	}

	mu.Unlock()
	w.Write([]byte(strconv.Itoa(counter)))
}
func parle(w http.ResponseWriter, r *http.Request) {
	fmt.Print("le btn a été cliqué")
}
