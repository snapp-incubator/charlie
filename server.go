package fwfag

import (
	"flag"
	"log"
	"net/http"
)

var (
	address = flag.String("listen", ":8080", "listen address")
	dir     = flag.String("dir", ".", "directory to serve")
)

func Run() {
	flag.Parse()

	log.Printf("listening on %q...", *address)
	log.Fatal(http.ListenAndServe(*address, http.FileServer(http.Dir(*dir))))
}
