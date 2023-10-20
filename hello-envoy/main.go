package main

import (
	"log"
	"net/http"
)

// not used any more
type Message struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//m := Message{Message: "Hello, Envoy!"}
		m := "Hello, Envoy!"
		/*b, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}*/
		//w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Type", "text/plain")
		//w.Write(b)
		w.Write(([]byte)(m))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
