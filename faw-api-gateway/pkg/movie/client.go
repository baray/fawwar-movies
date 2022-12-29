package movie

import (
	"fmt"

	"github.com/fawwar-movies/faw-api-gateway/pkg/config"
	"github.com/fawwar-movies/faw-api-gateway/pkg/movie/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.MovieServiceClient
}

func InitServiceClient(c *config.Config) pb.MovieServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.MovieSrvUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect to movies service:", err)
	}

	return pb.NewMovieServiceClient(cc)
}
