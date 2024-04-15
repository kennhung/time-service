package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := struct {
			DateTime string `json:"date_time"`
		}{time.Now().Format(time.RFC3339)}

		buf, err := json.Marshal(resp)
		if err != nil {
			log.Printf("failed to marshal response: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(buf)
		if err != nil {
			log.Printf("failed to write response: %v", err)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
		return
	}
}
