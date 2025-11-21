package main

import (
	"fmt"
	"net/http"
)


func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }







		
    }
}


//my own function that takes filename and serves the file now files all in one folder
func serveFile(w http.ResponseWriter, r *http.Request,filename string) {
	 ref := r.Header.Get("Referer")
    fmt.Println("Publisher URL:", ref)
		http.ServeFile(w, r, filename)
	}

func main() {

	// Serve tag.js
	http.HandleFunc("/server/tag.js", func(w http.ResponseWriter, r *http.Request) {
        serveFile(w, r, "tag1.js")
    })

	// Serve ad.html
	http.HandleFunc("/server/ad.html", func(w http.ResponseWriter, r *http.Request) {
		 serveFile(w, r, "ad.html")
	})

	// Serve ad.js
	http.HandleFunc("/server/ad.js", func(w http.ResponseWriter, r *http.Request) {
		 serveFile(w, r, "ad.js")
	})


	http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

	http.HandleFunc("/cross+origin+nestediframe/server/inneriframe.html", func(w http.ResponseWriter, r *http.Request) {
		 
		 serveFile(w, r, "inneriframe.html")
	})


	fmt.Println("🚀 Go Ad Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}








