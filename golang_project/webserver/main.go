package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World\n")
		io.WriteString(w, r.Method)
	}
	http.HandleFunc("/", h1)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
