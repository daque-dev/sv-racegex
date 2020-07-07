package api

import (
	"encoding/json"
	"log"
	"net/http"
	"racegex/database"
	"racegex/models"

	"github.com/gorilla/mux"
)

// GetLevels gets the list of all the levels
func GetLevels(w http.ResponseWriter, r *http.Request) {
	levels := []models.Level{}

	// Preload Lessons so the Lessons field is populated before returning
	res := database.DBConn.Preload("Lessons").Find(&levels)

	if res.Error != nil {
		log.Printf("Couldn't get levels")
	}

	if err := json.NewEncoder(w).Encode(res.Value); err != nil {
		log.Printf("Error parsing %-v", levels)
	}
}

// GetLevel gets a specific level in the list
func GetLevel(w http.ResponseWriter, r *http.Request) {
	// Get the url params of the route
	params := mux.Vars(r)

	level := models.Level{}

	res := database.DBConn.Preload("Lessons").Where("id = ?", params["id"]).First(&level)

	if res.Error != nil {
		log.Printf("Couldn't get level %-v", params["id"])
	}

	if err := json.NewEncoder(w).Encode(res.Value); err != nil {
		log.Printf("Error parsing %-v", level)
	}
}
