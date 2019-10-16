package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/callBack",callback)
	log.Fatal(http.ListenAndServe("107.173.26.195:8000", nil))
}

// handler echoes the Path component of the request URL r.
func callback(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	mid := query.Get("mid")
	signature := query.Get("signature")
	timestamp := query.Get("timestamp")
	nonce := query.Get("nonce")
	fmt.Fprintf(w,"GET: mid=%s </br> signature = %s </br> timestamp = %s </br> nonce = %s ", mid,signature,timestamp,nonce)
}