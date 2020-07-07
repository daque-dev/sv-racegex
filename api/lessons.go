package api

import (
	"encoding/json"
	"log"
	"net/http"
	"racegex/database"
	"racegex/models"

	"github.com/gorilla/mux"
)

// GetLessons gets the list of all the lessons
func GetLessons(w http.ResponseWriter, r *http.Request) {
	lessons := []models.Lesson{}

	res := database.DBConn.Find(&lessons)

	if res.Error != nil {
		log.Printf("Couldn't get lessons")
	}

	if err := json.NewEncoder(w).Encode(res.Value); err != nil {
		log.Printf("Error parsing %-v", lessons)
	}
}

// GetLesson gets a specific lesson in the list
func GetLesson(w http.ResponseWriter, r *http.Request) {
	// Get the url params of the route
	params := mux.Vars(r)

	lesson := models.Lesson{}

	res := database.DBConn.First(&models.Lesson{}, params["id"])

	if res.Error != nil {
		log.Printf("Couldn't get lesson %-v", params["id"])
	}

	if err := json.NewEncoder(w).Encode(res.Value); err != nil {
		log.Printf("Error parsing %-v", lesson)
	}
}
