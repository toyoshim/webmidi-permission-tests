package main

import (
  "fmt"
  "net/http"
  "os"
)

func main() {
  http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
  fmt.Println("listening...")
  err := http.ListenAndServe(":" + os.Getenv("PORT"), nil)
  if err != nil {
    panic(err)
  }
}
