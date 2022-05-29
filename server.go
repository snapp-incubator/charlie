package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	address = flag.String("listen", ":8080", "listen address")
	dir     = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()

	log.Printf("listening on %q...", *address)
	log.Fatal(http.ListenAndServe(*address, http.FileServer(http.Dir(*dir))))
}
