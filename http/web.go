package main

import (
    "log"
    "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Cool header</h1>"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", index)
    log.Fatal(http.ListenAndServe(":5000", mux))
}
