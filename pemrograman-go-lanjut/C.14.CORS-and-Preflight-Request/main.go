package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
    http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
        // w.Header().Set("Access-Control-Allow-Origin", "https://www.google.com, https://novalagung.com")
        // w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
        // w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

		w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "*")
        w.Header().Set("Access-Control-Allow-Headers", "*")

        if r.Method == "OPTIONS" {
            w.Write([]byte("allowed"))
            return
        }

		fmt.Println("MASUK URL INDEX")
        w.Write([]byte("hello dede"))
    })

    log.Println("Starting app at :9000")
    http.ListenAndServe(":9000", nil)
}