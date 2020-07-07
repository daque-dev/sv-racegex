package api

import (
	"encoding/json"
	"log"
	"net/http"
	"racegex/database"
	"racegex/models"

	"github.com/gorilla/mux"
)

// GetProblems gets the list of all the problems
func GetProblems(w http.ResponseWriter, r *http.Request) {
	problems := []models.Problem{}

	res := database.DBConn.Find(&problems)

	if res.Error != nil {
		log.Printf("Couldn't get problems")
	}

	if err := json.NewEncoder(w).Encode(res.Value); err != nil {
		log.Printf("Error parsing %-v", problems)
	}
}

// GetProblem gets a specific problem in the list
func GetProblem(w http.ResponseWriter, r *http.Request) {
	// Get the url params of the route
	params := mux.Vars(r)

	problem := models.Problem{}

	res := database.DBConn.First(&problem, params["id"])

	if res.Error != nil {
		log.Printf("Couldn't get problem %-v", params["id"])
	}

	if err := json.NewEncoder(w).Encode(res.Value); err != nil {
		log.Printf("Error parsing %-v", problem)
	}
}
