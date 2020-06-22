package database

import (
	"log"
	"racegex/config"

	"github.com/go-pg/pg/v10"
)

// GetConnection Create a connection with database, return this connection
func GetConnection() *pg.DB {
	var datastore config.Datastore
	datastore = config.GetDatabaseConfiguration()

	db := pg.Connect(&pg.Options{
		Addr:     datastore.Address + ":" + datastore.Port,
		User:     datastore.User,
		Password: datastore.Password,
		Database: datastore.Database,
	})

	log.Print("Successfully: Connection to database")

	return db
}

// CloseConnection Close the connection to the database
func CloseConnection(db *pg.DB) {
	db.Close()
	log.Print("Successfully: Close connection to database")
}
