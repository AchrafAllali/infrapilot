package main

import (
"encoding/json"
"log"
"net/http"
)

func main() {
mux := http.NewServeMux()

mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
})

mux.HandleFunc("GET /api/projects", func(w http.ResponseWriter, r *http.Request) {
json.NewEncoder(w).Encode([]map[string]string{{"name": "hello"}})
})

log.Println("InfraPilot écoute sur :8080")
log.Fatal(http.ListenAndServe(":8080", mux))
}
