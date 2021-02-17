package main

import (
    "log"
    "net/http"
    "time"
)

func index(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Cool header " + time.Now().Format("01-01-2006 15:04:05 Mon") + "</h1>"))
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://google.com", 301)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", index)
    mux.HandleFunc("/route", pageHandler)

    log.Fatal(http.ListenAndServe(":5000", mux))
}
