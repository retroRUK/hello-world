package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World Test App")
	})

	http.ListenAndServeTLS(":5000", os.Getenv("SSL_CERT"), os.Getenv("SSL_KEY"), nil)
}
