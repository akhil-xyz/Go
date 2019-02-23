package main

import (
	"io"
	"net/http"
)

func requestHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello World!")
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}
