package configuration

import (
	"gopkg.in/yaml.v3"
	"log"
)

type JobConfiguration struct {
	FirstJobConfiguration  FirstJob  `yaml:"first-job"`
	SecondJobConfiguration SecondJob `yaml:"second-job"`
}

type FirstJob struct {
	File File   `yaml:"file"`
	Host string `yaml:"host"`
}

type SecondJob struct {
	File File   `yaml:"file"`
	Host string `yaml:"host"`
}

type File struct {
	Name  string `yaml:"name"`
	Limit int    `yaml:"limit"`
}

var JobConfigurationVariable = &JobConfiguration{}

func LoadJobConfiguration(configData []byte) {
	err := yaml.Unmarshal(configData, JobConfigurationVariable)

	if err != nil {
		panic(err)
	}
	log.Printf("First-Job Configuration loaded from yaml, host: %s",
		JobConfigurationVariable.FirstJobConfiguration.Host)
	log.Printf("Second-Job Configuration loaded from yaml, host: %s",
		JobConfigurationVariable.SecondJobConfiguration.Host)
}
