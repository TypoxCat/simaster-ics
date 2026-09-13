// Modified from the original SimasterICSGen project to serve the web UI reliably.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	HTTP_PORT = 8080
)

func main() {
	// Read index.html once
	indexHTML, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Printf("Error reading index.html: %v\n", err)
		return
	}

	fs := http.FileServer(http.Dir("."))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Serve index.html for root path
		if r.URL.Path == "/" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(indexHTML)
			return
		}
		fs.ServeHTTP(w, r)
	})

	fmt.Printf("Listening on port %v\n", HTTP_PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", HTTP_PORT), nil)
}
