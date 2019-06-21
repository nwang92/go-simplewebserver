package main

import (
    "log"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	m := now.Minute()
	if (m%10 == 1 || m%10 == 2 || m%10 == 3 || m%10 == 5 || m%10 == 7 || m%10 == 8) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("☄ GOOD - 200 HTTP status code returned!"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("☄ BAD - 500 HTTP status code returned!"))
	}
	//w.WriteHeader(http.StatusOK)
	//w.Write([]byte("☄ GOOD - 200 HTTP status code returned!"))
}

func main() {
    http.HandleFunc("/simplewebserver/", handler)
    log.Fatal(http.ListenAndServe(":8001", nil))
}
