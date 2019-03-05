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

func swapScriptHandler(w http.ResponseWriter, r *http.Request) {
	mbInKb        := 1024
	req           := r.URL.Path[1:]
	swapSize, err := strconv.Atoi(req[:len(req) - 2])
	sizeType   	  := req[len(req)-2:]

	if err != nil {
		fmt.Fprintf(w, "")
	}

	if sizeType == "gb" || sizeType == "GB" {
		swapSize = swapSize * mbInKb
	}

	if sizeType == "tb" || sizeType == "TB" {
		swapSize = swapSize * mbInKb * mbInKb
	}

	script := `#!/bin/sh

echo 'Generating ` + r.URL.Path[1:] + ` file...'
sudo dd if=/dev/zero of=/swapfile bs=1M count=` + strconv.Itoa(swapSize) + `
echo 'Fixing permissions...'
sudo chmod 600 /swapfile
echo 'Setting generated file as swap...'
sudo mkswap /swapfile
sudo swapon /swapfile
echo 'Writing to fstab to make swap persistence...'
echo '/swapfile swap swap sw 0 0' | sudo tee --append /etc/fstab
echo 'Done!'`

	fmt.Fprintf(w, script)
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
