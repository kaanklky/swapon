package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Port - port number to run server on
var Port = 4040

// PortString - to use port number as in addr format ':xxxx'
var PortString = ":" + strconv.Itoa(Port)

func swapScriptHandler(w http.ResponseWriter, r *http.Request) {
	req := r.URL.Path[1:]
	swapSize := strings.ToUpper(req[:len(req)-1])

	script := `#!/bin/sh

sudo fallocate -l %s /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
echo '/swapfile swap swap sw 0 0' | sudo tee --append /etc/fstab`

	fmt.Fprintf(w, script, swapSize)
}

func main() {
	rtr := mux.NewRouter()
	http.Handle("/", rtr)

	rtr.Handle("/", http.FileServer(http.Dir("./web")))
	fmt.Printf("Running on http://127.0.0.1:%d/\n", Port)

	rtr.HandleFunc("/{size:[0-9]+(?:mb|MB|gb|GB|tb|TB)}", swapScriptHandler)
	fmt.Printf("Running on http://127.0.0.1:%d/{0-9}(mb|gb|tb)\n", Port)

	http.ListenAndServe(PortString, nil)
}
