package database

import (
	"log"
	"racegex/config"

	"github.com/jinzhu/gorm"
	// Postgres Dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// GetConnection Create a connection with database, return this connection
func GetConnection() *gorm.DB {
	var datastore config.Datastore = config.GetDatabaseConfiguration()

	db, err := gorm.Open("postgres", "host="+datastore.Address+" port="+datastore.Port+" user="+datastore.User+" dbname="+datastore.Database+" password="+datastore.Password)

	if err != nil {
		log.Fatalf("Error: Connection to database %v", err)
	}

	log.Print("Successfully: Connection to database")

	return db
}

// CloseConnection Close the connection to the database
func CloseConnection(db *gorm.DB) {
	db.Close()
	log.Print("Successfully: Close connection to database")
}
