package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(w io.Writer, name string) {
	_, err := fmt.Fprintf(w, "Hello, %v", name)
	if err != nil {
		return
	}
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
