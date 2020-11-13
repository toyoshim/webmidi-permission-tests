package main

import (
	"fmt"
	"net/http"
	"os"

//	"github.com/toyoshim/gomolog"
)

func main() {
//	log := gomolog.Open(os.Getenv("MONGOLAB_URI"), "log")
//	defer log.Close()

	fmt.Println("listening on " + os.Getenv("PORT") + " ...")
	http.Handle("/", http.FileServer(http.Dir("static")))
//	err := http.ListenAndServe(":"+os.Getenv("PORT"), log.Logger())
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
