package main

import (
	"log"
	"net/http"

	"github.com/joshadambell/brickandclick/routes"
	goji "goji.io"
)

func main() {
	mux := goji.NewMux()
	routes.Register(mux)
	go func() {
		err := http.ListenAndServe(":8081", mux)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("./assets")))
	if err != nil {
		log.Fatal(err)
	}
}
