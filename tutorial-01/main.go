package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.Int("port", 80, "port to run application")
	dir := flag.String("directory", "web/", "directory of html files")
	flag.Parse()

	fs := http.Dir(*dir)
	handler := http.FileServer(fs)
	http.Handle("/", handler)

	addr := fmt.Sprintf("localhost:%d", *port)
	fmt.Printf("Will listen on address %s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
