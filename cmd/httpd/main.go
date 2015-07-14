package main

import (
  "fmt"
  "net/http"
  "os"
)

func main() {
  fmt.Println("listening on " + os.Getenv("PORT") + " ...")
  err := http.ListenAndServe(":" + os.Getenv("PORT"), http.FileServer(http.Dir("static")))
  if err != nil {
    panic(err)
  }
}
