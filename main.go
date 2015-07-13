package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var port = flag.Int("port", 3000, "Port to run the server on")

var indexTemplate = template.Must(template.New("index").Parse(`<!DOCTYPE HTML>
<html>
  <head>
    <title>blog</title>
  </head>
  <body>
    <p>this is the beginning of the blog</p>
  </body>
</html>
`))

func main() {
	flag.Parse()

	http.HandleFunc("/", indexHandler)

	log.Printf("Starting server on port %d\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate.Execute(w, nil)
}
