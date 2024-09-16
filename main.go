package main

import (
	"job/configuration"
	"os"
)

func Init() {
	data, err := os.ReadFile("config.yaml")

	if err != nil {
		panic(err)
	}

	configuration.LoadJobConfiguration(data)
	configuration.LoadDatabaseConfiguration(data)
}

func main() {
	Init()
}
