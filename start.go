package main

import (
  "fmt"
  "strconv"
  "net/http"
  "pontius"
)

func ping(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong"))
}

func main() {
   pontius.New(http)
}
