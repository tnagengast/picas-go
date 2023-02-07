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

        fmt.Fprintf(w, "Hello World")
    })

    fmt.Println("Starting server at http://localhost:8080")

    http.ListenAndServe(":8080", nil)
}

