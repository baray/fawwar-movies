package watched

import (
	"github.com/fawwar-movies/faw-api-gateway/pkg/config"
	user "github.com/fawwar-movies/faw-api-gateway/pkg/user"
	"github.com/fawwar-movies/faw-api-gateway/pkg/watched/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, userSrv *user.ServiceClient) {
	a := user.InitUserMiddleware(userSrv)

	srv := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/watches")
	routes.Use(a.AuthRequired)
	routes.GET("/", srv.WatchedList)
	routes.POST("/:id/watch", srv.WatchMovie)
	routes.POST("/:id/rating", srv.RatingMovie)
}

func (srv *ServiceClient) WatchMovie(ctx *gin.Context) {
	handlers.WatchMovie(ctx, srv.Client)
}

func (srv *ServiceClient) RatingMovie(ctx *gin.Context) {
	handlers.RatingMovie(ctx, srv.Client)
}

func (srv *ServiceClient) WatchedList(ctx *gin.Context) {
	handlers.WatchedList(ctx, srv.Client)
}
