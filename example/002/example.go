package main

import (
	"io/ioutil"
	"net/http"

	"github.com/go-zoo/bonex"
)

func main() {
	mux := bonex.New()

	mux.Get("/", defaultHandler)
	mux.Get("/test", defaultHandler)
	mux.HandleFunc("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8080", mux)
}

func defaultHandler(rw http.ResponseWriter, req *http.Request, args bonex.Args) {
	file, _ := ioutil.ReadFile("index.html")
	rw.Write(file)
}
