package slurper

import (
	"io/ioutil"
	"l7-snake/configstruct"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

//Init just consumes a path to a file
func Init(file string) {
	if statFile(file) {
		log.Println("Config found! Using", file)
		configData, configError := ioutil.ReadFile(file)
		if configError != nil {
			log.Fatal("Could not read config file!", configError)
		}
		var rawYAML configstruct.Config
		ymlErr := yaml.Unmarshal(configData, &rawYAML)
		if ymlErr != nil {
			log.Fatal("Unmarshal error!", ymlErr)
		}
		configstruct.SetConfig(rawYAML)
		configstruct.SetConfigPath(file)
	}
}

func statFile(file string) bool {
	_, fileErr := os.Stat(file)
	if fileErr != nil {
		log.Println("Config not found. Creating template at ", file)
		configstruct.ConfigWriter(file)
	}
	return true
}
