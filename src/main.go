package main

import (
	"fmt"
	"net/http"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			myFigure := figure.NewFigure(r.URL.Path, "", true)
			fmt.Fprintf(w, myFigure.String())
			return
		}

		myFigure := figure.NewFigure("Hello! Welcome to Picas-Go :)", "", true)
		fmt.Fprintf(w, myFigure.String())
	})

	port := ":8081"
	fmt.Println("Listening on port", port)
	http.ListenAndServe(port, nil)
}
