package configchanger

import (
	"io/ioutil"
	"l7-snake/configstruct"
	"log"

	"gopkg.in/yaml.v3"
)

//this is a writing operation that flushes the current runtime data (CurrentConfig struct in configstruct.go) to the configfile
//everything is marshalled in yaml
func UpdateConfig() {
	//path to current config file
	cfgfile := configstruct.CurrentConfigPath

	//marshal interface to byte array (updated config)
	output, outputErr := yaml.Marshal(configstruct.CurrentConfig)
	if outputErr != nil {
		log.Fatal(outputErr)
	}
	//write updated data to file
	writeErr := ioutil.WriteFile(cfgfile, output, 0755)
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}
