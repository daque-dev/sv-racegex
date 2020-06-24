package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetProblems gets the list of all the problems
func GetProblems(w http.ResponseWriter, r *http.Request) {
	var data []Problem

	data = append(data, Problem{ID: "1", Title: "emails"})
	data = append(data, Problem{ID: "2", Title: "websites"})

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error parsing %-v", data)
	}
}

// GetProblem gets a specific problem in the list
func GetProblem(w http.ResponseWriter, r *http.Request) {
	// Get the url params of the route
	params := mux.Vars(r)

	// Initialize a mock array of items to look for the id
	var data []Problem

	data = append(data, Problem{ID: "1", Title: "emails"})
	data = append(data, Problem{ID: "2", Title: "websites"})

	for _, problem := range data {
		if problem.ID == params["id"] {
			if err := json.NewEncoder(w).Encode(problem); err != nil {
				log.Printf("Error parsing %-v", problem)
			}
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
