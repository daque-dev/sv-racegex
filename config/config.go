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
	Drop     bool
	Migrate  bool
}

// GetConfiguration Get configuration file
func GetConfiguration() (Configuration, error) {

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

	if err := json.Unmarshal(byteValue, &configuration); err != nil {
		return *(*Configuration)(nil), err
	}

	return configuration, nil
}

// GetDatabaseConfiguration Get database configuration
func GetDatabaseConfiguration() Datastore {
	var configuration Configuration
	configuration, err := GetConfiguration()
	if err != nil {
		log.Print(err)
		return *(*Datastore)(nil)
	}

	return configuration.Datastore
}
