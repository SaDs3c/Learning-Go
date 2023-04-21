package main

import (
	"fmt"
	"net/http"
)

const (
	host      = "localhost"
	port      = "8080"
	seperator = ":"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sayHello)

	fmt.Println("Listen on port", port)
	err := http.ListenAndServe(host+seperator+port, mux)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("Hello, World"))
		return
	}

	w.Write([]byte("error"))
	return
}
