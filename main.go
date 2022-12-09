package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Iniciamos web server en puerto 80")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola Mundo de Carolina")
	})
	http.ListenAndServe(":80", nil)
}
