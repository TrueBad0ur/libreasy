package main

import (
	"log"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/index.html")
}

func main() {
	http.HandleFunc("/", handleRoot)

	fs := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css", fs))

	fs = http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	fs = http.FileServer(http.Dir("./books"))
	http.Handle("/books/", http.StripPrefix("/books", fs))

	err := http.ListenAndServe("127.0.0.1:8081", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
