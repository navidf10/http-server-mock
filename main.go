package main

import (
	"net/http"
	"io"
)
var word string
func main () {
	http.HandleFunc("/", ping)
	http.ListenAndServe(":9090", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	txt := "0,SUCCESSFUL"
	io.WriteString(w, txt)
}