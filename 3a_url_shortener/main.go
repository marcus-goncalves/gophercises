package main

import (
	"fmt"
	"net/http"

	"example.com/v1/handler"
)

func main() {
	mux := defaultMux()

	// Build the map handler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handler.MapHandler(pathsToUrls, mux)

	// YAML handler using MapHandler as the fallback
	yml := `
- path: /urlshort
  url: https://godoc.org/gopkg.in/yaml.v2
- path: /urlshort-final
  url: https://github.com/gophercises/urlshortener/tree/solution
`
	yamlHandler, err := handler.YAMLHandler([]byte(yml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server :8080")
	http.ListenAndServe(":8080", yamlHandler)
	// http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
