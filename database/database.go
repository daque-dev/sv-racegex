// Package database connects and disconnects the main package to the
// postgres DB using gorm. The main variable available to any importing
// package is DBConn.
package database

import (
	"log"
	"racegex/config"
	"racegex/models"

	"github.com/jinzhu/gorm"
	// Postgres Dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// DBConn handles the connection to share between packages
	DBConn    *gorm.DB
	dbMigrate bool
	dbDrop    bool
)

// GetConnection creates a connection with the database using the
// configuration available at racegex/config/config.json
func GetConnection() {
	var datastore config.Datastore = config.GetDatabaseConfiguration()

	var err error
	DBConn, err = gorm.Open(
		"postgres",
		"sslmode=disable"+
			" host="+datastore.Address+
			" port="+datastore.Port+
			" user="+datastore.User+
			" dbname="+datastore.Database+
			" password="+datastore.Password)

	if err != nil {
		log.Fatalf("Error: Connection to database %v", err)
	}

	dbMigrate = datastore.Migrate
	dbDrop = datastore.Drop

	log.Print("Successfully: Connection to database")
}

// Migrate creates the necessary tables to fit the app types
func Migrate() {
	if !dbMigrate {
		log.Print("Won't migrate models")
		return
	}
	log.Print("Migrating models")
	migrate(&models.Problem{})
	migrate(&models.Lesson{})
}

// Delete existing table and automigrate from models
func migrate(table interface{}) {
	log.Printf(" - %T", table)

	if dbDrop {
		if err := DBConn.DropTableIfExists(table).Error; err != nil {
			log.Print("    - Failed drop")
		} else {
			log.Print("    - Successful drop")
		}
	}
	if err := DBConn.AutoMigrate(table).Error; err != nil {
		log.Printf("    - Failed migration")
	} else {
		log.Printf("    - Successful migration")
	}
}

// CloseConnection closes the connection to the database
func CloseConnection() {
	if err := DBConn.Close(); err != nil {
		log.Print("Failure: Close connection to database")
	}
	log.Print("Successfully: Close connection to database")
}
