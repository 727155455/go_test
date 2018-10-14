package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func sayHelloNmae(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("path", r.URL.Path)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "xxxxx, %d\n", 123)
	data, err := json.Marshal(r)
	if err != nil {
		fmt.Println("xxxxxx", err)
	}

	fmt.Fprintln(w, string(data))
	fmt.Printf("%+v\n", r)
	fmt.Printf("%+v\n,%#v\n,%T\n", r.Context(), r.Context(), r.Context())
}

func main() {
	http.HandleFunc("/", sayHelloNmae)
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/foo2", func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(responseWriter, "Hi, %q", html.EscapeString(request.URL.Path))
	})
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("xxxx", err)
	}
}