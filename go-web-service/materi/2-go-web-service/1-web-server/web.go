package main

import (
	"fmt"
	"net/http"
)

// PORT untuk server
var PORT = ":8080"

// Fungsi untuk menghandle response
func greet (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	/**
	 * Web server pada go
	 */

	/**
	 * HanldeFunc adalah fungsi untuk mengatur routing yang menerima dua parameter
	 * parameter pertama adalah route-nya
	 * parameter kedua adalah fungsi dengan 2 parameter (http.ResponseWriter, *http.Request)
	 */
	http.HandleFunc("/", greet)

	http.HandleFunc("/kor", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "안녕")
	})

    http.ListenAndServe(PORT, nil)
}