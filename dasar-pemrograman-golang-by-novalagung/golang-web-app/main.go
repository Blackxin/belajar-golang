package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	address := ":9000"
  fmt.Printf("Server started at http://localhost:%s\n", address)
	/* err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	} */
	server := new(http.Server)
	server.Addr = address
  server.ReadTimeout = 10 * time.Second
  server.WriteTimeout = 10 * time.Second
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	message := "Welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	message := "Hello world"
	w.Write([]byte(message))
}
