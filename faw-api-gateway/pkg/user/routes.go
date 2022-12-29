package user

import (
	"github.com/gin-gonic/gin"
	"github.com/fawwar-movies/faw-api-gateway/pkg/user/handlers"
	"github.com/fawwar-movies/faw-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	srv := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", srv.Register)
	routes.POST("/login", srv.Login)

	return srv
}

func (srv *ServiceClient) Register(ctx *gin.Context) {
	handlers.Register(ctx, srv.Client)
}

func (srv *ServiceClient) Login(ctx *gin.Context) {
	handlers.Login(ctx, srv.Client)
}
