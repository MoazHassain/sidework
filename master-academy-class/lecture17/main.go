package main

import (
	"fmt"
	"net/http"
)

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }
func main() {

	http.HandleFunc("/", home)
	http.ListenAndServe(":8081", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `welcome to my first golang webpage`)
}
