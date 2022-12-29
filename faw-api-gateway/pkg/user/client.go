package user

import (
	"fmt"

	"github.com/fawwar-movies/faw-api-gateway/pkg/user/pb"
	"github.com/fawwar-movies/faw-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.UserServiceClient
}

func InitServiceClient(c *config.Config) pb.UserServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.UserSrvUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect to user service:", err)
	}

	return pb.NewUserServiceClient(cc)
}
