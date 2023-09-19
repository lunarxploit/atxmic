package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<h1>Atxmic is running on :8080</h1>`)
}

func serveCss(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "client/game.css")
}
