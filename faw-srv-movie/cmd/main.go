package main

import (
	"fmt"
	"log"
	"net"

	"github.com/fawwar-movies/faw-srv-movie/pkg/config"
	"github.com/fawwar-movies/faw-srv-movie/pkg/db"
	pb "github.com/fawwar-movies/faw-srv-movie/pkg/pb"
	services "github.com/fawwar-movies/faw-srv-movie/pkg/services"
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

	fmt.Println("FAWWAR Movies Srv on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterMovieServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
