package main

import (
	"fmt"
	"log"
	"os"

	micro "github.com/micro/go-micro"
	pb "github.com/ruprict/evalentine/consignment-service/proto/consignment"
	vesselProto "github.com/ruprict/evalentine/vessel-service/proto/vessel"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	defer session.Close()

	if e != nil {

		log.Panicf("Could not connect to DB with host %s - %v", host, err)
	}

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
