package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	A int `json:"a"`
	B int `json:"b"`
	C int `json:"c"`
}

type Response struct {
	Result int `json:"result"`
}

func middleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Result: 10})
		return
	}

	if r.Method == http.MethodPost {
		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		middle := findMiddle(req.A, req.B, req.C)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Result: middle})
		return
	}

	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}

func findMiddle(a, b, c int) int {
	if (a > b && a < c) || (a > c && a < b) {
		return a
	}
	if (b > a && b < c) || (b > c && b < a) {
		return b
	}
	return c
}

func main() {
	http.HandleFunc("/middle", middleHandler)

	addr := "127.0.0.1:8080"
	log.Printf("Server is running at %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
