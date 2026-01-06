package main

import (
	"log"
	"net/http"

	publicH "example.com/pkg/public/handler"

	"example.com/pkg/handler"
)

func main() {
	mux := http.NewServeMux()

	// Templates
	mux.HandleFunc("/hello-world", handler.HelloWorldPage)

	// Redirects requests from "/" to the landing page
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/landing", http.StatusFound)
	})

	mux.HandleFunc("/landing", publicH.LandingPage)
	mux.HandleFunc("/about_us", publicH.AboutUsPage)
	mux.HandleFunc("/blog", publicH.BlogPage)
	mux.HandleFunc("/contact", publicH.ContactPage)

 	if err := http.ListenAndServe(":8090", mux); err != nil {
		log.Fatal(err)
	}
}
