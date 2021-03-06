package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	api "racegex/api"
	"racegex/database"
	"racegex/seeds"
	socket "racegex/socket"
)

func main() {
	// This creates the database.DBConn that is accessible by all packages
	database.GetConnection()
	// Create the Tables for the racegex/models
	database.Migrate()
	// Seed the database
	seeds.Seed()
	// Close the connection when the main function returns
	defer database.CloseConnection()

	// Create a gorilla/mux router
	r := mux.NewRouter()

	r.HandleFunc("/problems", api.GetProblems)
	r.HandleFunc("/problems/{id}", api.GetProblem)
	r.HandleFunc("/lessons", api.GetLessons)
	r.HandleFunc("/lessons/{id}", api.GetLesson)

	// Create and start the WebSocket Hub
	hub := socket.NewHub()
	go hub.Run()
	// Specify the route for the WebSocket
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		socket.ServeWs(hub, w, r)
	})

	// Start the server
	log.Println(fmt.Sprintf("Server running on http://localhost%s", ":4000"))
	err := http.ListenAndServe(":4000", cors.Default().Handler(r))
	if err != nil {
		log.Panicf("could not run the server %v", err)
		return
	}
}
