package watched

import (
	"fmt"

	"github.com/fawwar-movies/faw-api-gateway/pkg/config"
	"github.com/fawwar-movies/faw-api-gateway/pkg/watched/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.WatchedServiceClient
}

func InitServiceClient(c *config.Config) pb.WatchedServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.WatchedSrvUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewWatchedServiceClient(cc)
}
