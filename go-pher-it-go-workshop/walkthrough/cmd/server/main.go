package main

import (
	"log"
	"net/http"

	"github.com/joerdav/myapp/api/hello"
)

func main() {
	m := http.NewServeMux()
	m.Handle("/hello", hello.Handler{})
	log.Println("Server started.")
	http.ListenAndServe("localhost:4444", m)
}
