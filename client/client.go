package client

import (
	"context"
	"l7-snake/centraldata"
	"l7-snake/configstruct"
	"l7-snake/pt3status"
	"log"
	"time"

	"google.golang.org/grpc"
)

//health indicator
var ConnPool int

func StartClient() {
	//the length equals the amount of targets we have
	looplen := len(configstruct.CurrentConfig.Data.Communication.Targets)

	//parsing delay interval for requests
	delay, _ := time.ParseDuration(configstruct.CurrentConfig.Data.Settings.Interval) //we already checked this exact value, no need to do it again

	//infinite loop for functions to call
	for {
		//wipe old data
		ConnPool = 0
		centraldata.CentralData = nil
		centraldata.CentralData = append(centraldata.CentralData, centraldata.OwnData())

		//call all targets
		for i := 0; i < looplen; i++ {
			connections(configstruct.CurrentConfig.Data.Communication.Targets[i])
		}

		//set health after connection attempts
		setHealth()

		//print centraldata (human readability)
		for i := 0; i < len(centraldata.CentralData); i++ {
			health := centraldata.CentralData[i].Health
			id := centraldata.CentralData[i].Id
			term := centraldata.CentralData[i].Terminator
			up := centraldata.CentralData[i].LastUpdated
			tr := centraldata.CentralData[i].Targets
			r := centraldata.CentralData[i].Routes
			if term == "yes" {
				log.Printf("%s\t| %s\t| %s\t| (-/-)\t| %s | Routes: %s", health, id, term, up, r)
			} else {
				log.Printf("%s\t| %s\t| %s\t| (%d/%d)\t| %s | Routes: %s", health, id, term, ConnPool, tr, up, r)
			}
		}

		time.Sleep(delay) //delay between calls
	}
}

func connections(ip string) {
	//connection object
	var conn *grpc.ClientConn

	//grpc dial to ip
	conn, cErr := grpc.Dial(ip, grpc.WithInsecure())
	if cErr != nil {
		log.Println("Client cannot dial to", ip)
		log.Fatal(cErr)
	}
	defer conn.Close() //default precaution

	//register new client service on connection object
	rq := pt3status.NewEchoServiceClient(conn)

	//Sssend the snake for echo
	msg := pt3status.Echo{
		Echo: "sss",
	}

	//rpc call context with timeout per call
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	//get rpc server response
	//ret, retErr := rq.Poke(context.Background(), &msg)
	ret, retErr := rq.Poke(ctx, &msg)
	if retErr != nil {

	}

	//adjust centraldata after received dataset
	if ret != nil {
		checkEntry(ret)
		ConnPool += 1
	}
}

//check and adjust centraldata if new data is gathered
func checkEntry(ret *pt3status.Status) {
	//looplen for received data
	l2 := len(ret.Lst)
	for i := 0; i < l2; i++ {
		centraldata.CentralData = append(centraldata.CentralData, ret.Lst[i])
	}
}

func setHealth() {
	//Healthy responds from all targets
	if ConnPool == len(configstruct.CurrentConfig.Data.Communication.Targets) {
		centraldata.CentralData[0].Health = "HEALTHY"
	}
	//No responds from targets
	if ConnPool == 0 {
		centraldata.CentralData[0].Health = "OFFLINE"
	}
	//Some responses from targets
	if ConnPool > 0 && ConnPool < len(configstruct.CurrentConfig.Data.Communication.Targets) {
		centraldata.CentralData[0].Health = "UNHEALTHY"
	}
}
