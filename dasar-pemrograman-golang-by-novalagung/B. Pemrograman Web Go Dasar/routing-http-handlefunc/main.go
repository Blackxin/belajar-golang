package main

import (
	"fmt"
	"net/http"
)

func main() {
  // penggunaan http.HandleFunc
  handlerIndex := func (w http.ResponseWriter, r *http.Request)  {
    w.Write([]byte("Hello"))
  }

  http.HandleFunc("/", handlerIndex)
  http.HandleFunc("/index", handlerIndex)

  http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello again"))
  })

  fmt.Println("Server started at http://localhost:9000")
  http.ListenAndServe(":9000", nil)
}
