package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/b/api", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		writer.Write([]byte("debug project b"))
	})
	http.ListenAndServe(":8000", nil)
}
