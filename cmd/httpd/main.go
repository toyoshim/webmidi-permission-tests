package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/toyoshim/gomolog"
)

func main() {
	log := gomolog.Open(os.Getenv("MONGOLAB_URI"), "log")
	defer log.Close()

	fmt.Println("listening on " + os.Getenv("PORT") + " ...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), http.FileServer(http.Dir("static")), log.Logger())
	if err != nil {
		panic(err)
	}
}
