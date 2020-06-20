package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/problems", getProblems)
	mux.HandleFunc("/problem", getProblem)

	log.Println(fmt.Sprintf("Server running on http://localhost%s", ":4000"))
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatalf("could not run the server %v", err)
		return
	}
}

func getProblems(w http.ResponseWriter, r *http.Request) {
	var data []Problem

	data = append(data, Problem{ID: "1", Title: "emails"})
	data = append(data, Problem{ID: "2", Title: "websites"})

	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getProblem(w http.ResponseWriter, r *http.Request) {
	data := Problem{ID: "1", Title: "emails"}

	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Problem : Contains a problem to challenge the user with
type Problem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
