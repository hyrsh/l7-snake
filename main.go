package main

import (
	"flag"
	"l7-snake/checkconfig"
	"l7-snake/client"
	"l7-snake/configstruct"
	"l7-snake/server"
	"l7-snake/slurper"
)

//I build with
//$env:CGO_ENABLED = 1; go build -ldflags='-s -w -extldflags "-static"' main.go
//packed with
//upx --best l7-snake.exe
//brute compression throws a false positive with Windows Defender -.-"

func init() {
	//cli flags for config file path
	configPath := flag.String("config", "./config.yml", "Path to config file")
	flag.Parse()

	//read config at path or create a default template
	slurper.Init(*configPath)

	//checkconfig
	checkconfig.Init()
}

func main() {
	if !configstruct.CurrentConfig.Data.Routing.Terminator {
		go client.StartClient()
	}
	server.StartServer()
}
