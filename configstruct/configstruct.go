package configstruct

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

//CurrentConfig stores our config at runtime (can also be updated during runtime)
var CurrentConfig Config

//CurrentConfigPath ... pretty self-explanatory
var CurrentConfigPath string

//Config sets the pattern for our config
type Config struct {
	Data struct {
		Communication struct {
			Id         string   `yaml:"id"`
			ListenPort string   `yaml:"listenport"`
			Targets    []string `yaml:"targets"`
		} `yaml:"communication"`
		Routing struct {
			Routes     []string `yaml:"routes"`
			Terminator bool     `yaml:"terminator"`
		} `yaml:"routing"`
		Settings struct {
			Interval string `yaml:"interval"`
		} `yaml:"settings"`
	} `yaml:"data"`
}

//SetConfig sets the config for central access (and partial updating (port changes still require a restart))
func SetConfig(config Config) {
	CurrentConfig = config
}

//SetConfigPath ... and again
func SetConfigPath(configpath string) {
	CurrentConfigPath = configpath
}

//ConfigWriter only gets called when no config is found to provide a template
func ConfigWriter(file string) {
	template := Config{}
	//set default values
	writeDefaults(&template)
	//marshal interface to byte array
	output, outputErr := yaml.Marshal(&template)
	if outputErr != nil {
		log.Println("YAML marshal error!")
		log.Fatal(outputErr)
	}
	//make sure the path exists and then write template to file
	filePath := filepath.Dir(file)
	pErr := os.MkdirAll(filePath, 0755)
	if pErr != nil {
		log.Println("Cannot create path", filePath)
		log.Fatal(pErr)
	}
	//write data to file
	writeErr := ioutil.WriteFile(file, output, 0755)
	if writeErr != nil {
		log.Println("YAML cannot write data to file!")
		log.Fatal(writeErr)
	}
}

//kind of self-explanatory. We set default values since this stupid struct is not able to do this on its own
func writeDefaults(config *Config) {
	//Communication
	config.Data.Communication.Id = "none"
	config.Data.Communication.ListenPort = "9001"
	config.Data.Communication.Targets = []string{"0.0.0.0:9002", "0.0.0.0:9003"}
	//Routing
	config.Data.Routing.Routes = []string{"route-a", "route-b"}
	config.Data.Routing.Terminator = true
	//Settings
	config.Data.Settings.Interval = "500ms"
}
