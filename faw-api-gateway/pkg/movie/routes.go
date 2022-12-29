package movie

import (
	"github.com/fawwar-movies/faw-api-gateway/pkg/config"
	"github.com/fawwar-movies/faw-api-gateway/pkg/movie/handlers"
	"github.com/fawwar-movies/faw-api-gateway/pkg/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, userSrv *user.ServiceClient) {
	a := user.InitUserMiddleware(userSrv)

	srv := &ServiceClient{
		Client: InitServiceClient(c),
	}

	r.GET("movie-list", srv.MoviesList)
	r.Static("/covers", "./media")

	routes := r.Group("/movies")
	routes.Use(a.AuthRequired)
	routes.POST("/", srv.CreateMovie)
	routes.GET("/:id", srv.FindMovie)
	routes.DELETE("/:id", srv.DeleteMovie)
	routes.PUT("/:id", srv.UpdateMovie)
	routes.POST("/:id/update-cover", srv.UploadCover)
}

func (srv *ServiceClient) CreateMovie(ctx *gin.Context) {
	handlers.CreateMovie(ctx, srv.Client)
}

func (srv *ServiceClient) FindMovie(ctx *gin.Context) {
	handlers.FindMovie(ctx, srv.Client)
}

func (srv *ServiceClient) DeleteMovie(ctx *gin.Context) {
	handlers.DeleteMovie(ctx, srv.Client)
}

func (srv *ServiceClient) UpdateMovie(ctx *gin.Context) {
	handlers.UpdateMovie(ctx, srv.Client)
}

func (srv *ServiceClient) UploadCover(ctx *gin.Context) {
	handlers.UploadCover(ctx, srv.Client)
}

func (srv *ServiceClient) MoviesList(ctx *gin.Context) {
	handlers.MoviesList(ctx, srv.Client)
}
