package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	hub := newHub()
	go hub.run()

	r.HandleFunc("/problems", getProblems).Methods("GET")
	r.HandleFunc("/problems/{id}", getProblem).Methods("GET")

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	log.Println(fmt.Sprintf("Server running on http://localhost%s", ":4000"))
	err := http.ListenAndServe(":4000", r)
	if err != nil {
		log.Fatalf("could not run the server %v", err)
		return
	}
}

func getProblems(w http.ResponseWriter, r *http.Request) {
	var data []Problem

	data = append(data, Problem{ID: "1", Title: "emails"})
	data = append(data, Problem{ID: "2", Title: "websites"})

	json.NewEncoder(w).Encode(data)
}

func getProblem(w http.ResponseWriter, r *http.Request) {
	// Get the url params of the route
	params := mux.Vars(r)

	// Initialize a mock array of items to look for the id
	var data []Problem

	data = append(data, Problem{ID: "1", Title: "emails"})
	data = append(data, Problem{ID: "2", Title: "websites"})

	for _, problem := range data {
		if problem.ID == params["id"] {
			json.NewEncoder(w).Encode(problem)
			return
		}
	}
	json.NewEncoder(w).Encode(Problem{})
}

// Problem : Contains a problem to challenge the user with
type Problem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
