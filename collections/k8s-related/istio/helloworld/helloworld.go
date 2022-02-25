package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	log.Println("Server hello-world")
	http.HandleFunc("/", AppRouter)
	http.ListenAndServe(":12345", nil)
}

func AppRouter(w http.ResponseWriter, r *http.Request) {
	dump, _ := httputil.DumpRequest(r, false)
	log.Printf("%q\n", dump)
	io.WriteString(w, "HELLO WORLD1111\n")
	return
}
