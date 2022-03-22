package main

import (
	"fmt"
	"net/http"

	"rsc.io/quote"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, quote.Hello())
}

func main() {
	http.HandleFunc("/", hello)
        fmt.Printf("Launching server on port 8081")
	http.ListenAndServe(":8081", nil)
}
