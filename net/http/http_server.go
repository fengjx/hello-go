package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("pong"))
		if err != nil {
			log.Fatalln(err.Error())
		}
	})

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
