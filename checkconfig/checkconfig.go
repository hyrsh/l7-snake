package checkconfig

import (
	"io/ioutil"
	"l7-snake/configstruct"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

//Init the settings check
func Init() {
	//check ID
	checkID(configstruct.CurrentConfig.Data.Communication.Id)

	//check listen port for validity
	checkPort(configstruct.CurrentConfig.Data.Communication.ListenPort)

	//check target IPs for validity
	l := len(configstruct.CurrentConfig.Data.Communication.Targets)
	for i := 0; i < l; i++ {
		div := strings.Split(configstruct.CurrentConfig.Data.Communication.Targets[i], ":")
		checkIP(div[0])
		checkPort(div[1])
	}

	//check terminator
	checkBool(configstruct.CurrentConfig.Data.Routing.Terminator)

	//check time
	checkTime(configstruct.CurrentConfig.Data.Settings.Interval)

	//Success message
	log.Println("Config ready for launch!")
}

//check ID for none or empty and generate + write a random one
func checkID(id string) {
	const keylen int = 20
	dummyID := ""
	rand.Seed(time.Now().UnixNano())
	pool := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-"}
	if id == "" || id == "none" {
		for i := 0; i < keylen; i++ {
			p := rand.Intn(len(pool))
			dummyID += pool[p]
		}

		//update ID in current config
		configstruct.CurrentConfig.Data.Communication.Id = dummyID

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
}

//check given string after conversion to uint16 for validity of range and exceptions
func checkPort(portStr string) {
	//we want to restrict some port allocations
	portArray := [4]uint16{0, 22, 80, 443}              //no 0, no SSH, no HTTP, no HTTPS
	port, portErr := strconv.ParseUint(portStr, 10, 16) //parse for uint16 decimal (base 10) representation, since the port range is numerically based on it
	if portErr != nil {
		log.Println("Port is invalid", portStr)
		log.Fatal(portErr)
	}
	for _, v := range portArray {
		if uint16(port) == v { //type casting is allowed because of strconv.ParseUint
			log.Fatal("Port " + portStr + " is not allowed to set!")
		}
	}
}

//check IP address validity
func checkIP(ipadr string) {
	ip := net.ParseIP(ipadr)
	if ip == nil {
		log.Fatal("IP not valid (whitespaces?) " + ipadr)
	}
}

//check boolean statement in the terminator field
func checkBool(state bool) {
	if state {
		//it is yes or true
		log.Println("The terminator flag is active, no targets will be gathered.")
	} else {
		//it is no or false
	}
}

func checkTime(t string) {
	_, tErr := time.ParseDuration(t)
	if tErr != nil {
		log.Println("Time not valid (format: ms, s, m or h)!")
		log.Println("e.g.: 500ms, 2s, 33m, 10h")
		log.Fatal("Given time " + t)
	}
}
