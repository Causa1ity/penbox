package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// Run Startup program listening to the specified address
func Run(addr string) {
	s := Status{
		"penbox running...",
		make(map[string]string),
	}

	http.HandleFunc("/payloads/flask/session", flaskSession)
	s.Functions["flaskSession"] = "/payloads/flask/session"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/json")
		status, _ := json.Marshal(s)
		w.Write(status)
	})
	log.Fatal(http.ListenAndServe(addr, nil))
}

type Status struct {
	ServeStatus string            `json:"status"`
	Functions   map[string]string `json:"functions"`
}
