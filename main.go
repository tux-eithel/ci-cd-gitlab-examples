package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := flag.Int("p", 8080, "listen port")
	flag.Parse()

	http.HandleFunc("/", hello)

	fmt.Printf("Listen new connection at http://localhost:%d\n", *port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatalf("Unable to run server: %v", err)
	}

}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, sayHi(r.URL.Query().Get("name")))
}

// sayHi is the entry-point for our web application
func sayHi(s string) string {
	hello := ""
	if s != "" {
		hello += fmt.Sprintf("Hi %s!\n", s)
	}
	hello += "Welcome to our website"
	return hello
}
