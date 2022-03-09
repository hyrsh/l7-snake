package server

import (
	"context"
	"l7-snake/centraldata"
	"l7-snake/configstruct"
	"l7-snake/pt3status"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

//empty server struct for grpc
type Server struct {
}

func (s *Server) Poke(ctx context.Context, message *pt3status.Echo) (*pt3status.Status, error) {
	centraldata.CentralData[0].LastUpdated = time.Now().Format("2006-01-02 15:04:05")

	return &pt3status.Status{Lst: centraldata.CentralData}, nil
}

//start grpc server
func StartServer() {
	//initial entry
	if len(centraldata.CentralData) == 0 {
		centraldata.CentralData = append(centraldata.CentralData, centraldata.OwnData())
	}
	//endpoint health is always "HEALTHY" since an "OFFLINE" state can only be reported by an external client requesting
	//something similar applies to the targets (it does not have any), which makes zero a special number
	if configstruct.CurrentConfig.Data.Routing.Terminator {
		centraldata.CentralData[0].Health = "HEALTHY"
		centraldata.CentralData[0].Targets = 0
		centraldata.CentralData[0].Routes[0] = "END OF " + configstruct.CurrentConfig.Data.Routing.Routes[0] //we force one endpoint to have one ending route
	}

	//get port from slurped config
	port := ":" + configstruct.CurrentConfig.Data.Communication.ListenPort

	//default net listener
	listener, lErr := net.Listen("tcp", port)
	if lErr != nil {
		log.Fatal("Could not start listening!", lErr)
	}
	defer listener.Close() //default precaution

	//grpc server listen
	g := grpc.NewServer()

	//register grpc
	pt3status.RegisterEchoServiceServer(g, &Server{})

	//echo debug
	log.Println("gRPC server listening on", configstruct.CurrentConfig.Data.Communication.ListenPort)
	log.Println("My ID:", configstruct.CurrentConfig.Data.Communication.Id)

	if gErr := g.Serve(listener); gErr != nil {
		log.Fatal("Could not start gRPC listening!", gErr)
	}
}
