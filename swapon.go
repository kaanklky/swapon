package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Port - port number to run server on
var Port = 4040

// PortString - to use port number as in addr format ':xxxx'
var PortString = ":" + strconv.Itoa(Port)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
}

func swapScriptHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("swapfile"))
}

func main() {
	rtr := mux.NewRouter()
	http.Handle("/", rtr)

	rtr.HandleFunc("/", indexHandler)
	fmt.Printf("Running on http://127.0.0.1:%d/\n", Port)

	rtr.HandleFunc("/{size:[0-9]+(?:mb|gb|tb)}", swapScriptHandler)
	fmt.Printf("Running on http://127.0.0.1:%d/{0-9}(mb|gb|tb)\n", Port)

	http.ListenAndServe(PortString, nil)
}
