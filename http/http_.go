package main

import (
  "encoding/json"
  "net/http"
)

type Payload struct {
  Version    float32
  Content []string
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  payload := Payload{1.0, []string{"Amazingly", "Surprisingly", "Ridiculously"}}

  content, err := json.Marshal(payload)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(content)
}