package main

import (
	"fmt"
	"log"
	"net"

	"github.com/fawwar-movies/faw-srv-movie/pkg/db"
	"github.com/fawwar-movies/faw-srv-watched/pkg/config"
	pb "github.com/fawwar-movies/faw-srv-watched/pkg/pb"
	services "github.com/fawwar-movies/faw-srv-watched/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("FAWWAR Watches Srv on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterWatchedServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
