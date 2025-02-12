package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)

	println("start server")

	err := http.ListenAndServe(":8099", nil)
	if err != nil {
		panic(err)
	}

}
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world"))
}
