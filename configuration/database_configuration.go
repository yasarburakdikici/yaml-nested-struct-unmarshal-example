package configuration

import (
	"gopkg.in/yaml.v3"
	"log"
)

type DatabaseConfiguration struct {
	Database Database `yaml:"database"`
}

type Database struct {
	Username string `yaml:"host"`
	Password string `yaml:"password"`
	Url      string `yaml:"url"`
	Schema   string `yaml:"schema"`
}

var DatabaseConfigurationVariable = &DatabaseConfiguration{}

func LoadDatabaseConfiguration(configData []byte) {
	err := yaml.Unmarshal(configData, DatabaseConfigurationVariable)

	if err != nil {
		panic(err)
	}
	log.Printf("Database Configuration loaded from yaml, db-url: %s",
		DatabaseConfigurationVariable.Database.Url)
}
