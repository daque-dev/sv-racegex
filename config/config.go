package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Configuration Struct for configuration server
type Configuration struct {
	Datastore Datastore `json:"datastore"`
}

// Datastore Struct for configuration database
type Datastore struct {
	Address  string `json:"address"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// GetConfiguration Get configuration file
func GetConfiguration() Configuration {

	file, err := os.Open("config/config.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatalf("Not found configuration file %v", err)
	}

	log.Println("Successfully: Opened configuration")

	// defer the closing of our file so that we can parse it later on
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var configuration Configuration

	json.Unmarshal(byteValue, &configuration)

	return configuration
}

// GetDatabaseConfiguration Get database configuration
func GetDatabaseConfiguration() Datastore {
	var configuration Configuration

	configuration = GetConfiguration()

	return configuration.Datastore
}
