package main

import (
	"fmt"
	"net/http"
)

func main() {
	hello := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World from net/http server")
	}

	http.HandleFunc("/", hello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
