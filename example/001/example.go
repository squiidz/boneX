package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-zoo/bonex"
)

func main() {
	mux := bonex.New()

	// Eval takes func(string) bool as arguments,
	// eval bind the first function with the first URL params, the second function
	// with the second params, etc ...
	mux.Get("/index/:test/:text", varHandler).Eval(isANumber, biggerThen2)

	mux.Post("/data", dataHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func varHandler(rw http.ResponseWriter, req *http.Request, args bonex.Args) {
	test := args.GetValue("test")
	text := args.GetValue("text")

	log.Println("TEST =", test)
	log.Println("TEXT =", text)

	rw.Write([]byte(args.GetValue("test")))
}

func dataHandler(rw http.ResponseWriter, req *http.Request, args bonex.Args) {
	rw.Write([]byte("Some useless data ..."))
}

// Evaluator which check if the url parameters is a number
func isANumber(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return true
	}
	return false
}

// Evaluator which check if the url parameters is bigger than 2
func biggerThen2(str string) bool {
	if len(str) < 2 {
		return true
	}
	return false
}
