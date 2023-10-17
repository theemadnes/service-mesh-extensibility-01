package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m := Message{Message: "Hello, Envoy!"}
		b, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
