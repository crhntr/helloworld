package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		str := html.EscapeString(r.URL.Path)
		if len(str) > 0 {
			fmt.Fprintf(w, "Hello, %s!", str[1:])
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
