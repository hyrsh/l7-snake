package main

import (
	"flag"
	"l7-snake/banner"
	"l7-snake/checkconfig"
	"l7-snake/client"
	"l7-snake/configstruct"
	"l7-snake/server"
	"l7-snake/slurper"
	"log"
)

//On Windows
//$env:CGO_ENABLED = 1; go build -ldflags='-s -w -extldflags "-static"' main.go

//On Linux (important for docker image building with "FROM scratch")
//go build -a -tags netgo --ldflags '-extldflags "-static"'

//packed with
//upx --best l7-snake.exe (Windows)
//or
//upx --best l7-snake (Linux)
//brute compression throws a false positive with Windows Defender -.-"

func init() {
	//print banner
	log.Println(banner.PrintBanner())

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
