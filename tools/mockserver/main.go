package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/cases", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		fmt.Printf("POST /cases: Content-Type=%s, body_size=%d\n",
			r.Header.Get("Content-Type"), len(body))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"id":"abc-123","status":"processing"}`))
	})

	http.HandleFunc("/cases/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)

		w.Header().Set("Content-Type", "application/json")

		rest := strings.TrimPrefix(r.URL.Path, "/cases/")

		// /cases/{id}/status
		if strings.HasSuffix(rest, "/status") {
			id := strings.TrimSuffix(rest, "/status")
			if strings.HasPrefix(id, "missing") {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `{"error":"case %s not found"}`, id)
				return
			}
			fmt.Fprintf(w, `{"id":"%s","status":"done"}`, id)
			return
		}

		// /cases/{id}
		id := rest
		if strings.HasPrefix(id, "missing") {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, `{"error":"case %s not found"}`, id)
			return
		}
		fmt.Fprintf(w, `{"id":"%s","status":"done","alert":{"signature":"test"},"events":[]}`, id)
	})

	fmt.Println("mock backend on :8000")
	http.ListenAndServe(":8000", nil)
}
