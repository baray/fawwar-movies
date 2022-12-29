package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/fawwar-movies/faw-api-gateway/pkg/watched/pb"
	"github.com/gin-gonic/gin"
)

func WatchMovie(ctx *gin.Context, c pb.WatchedServiceClient) {
	movieId, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	userId, _ := ctx.Get("currentUserId")

	res, err := c.WatchMovie(context.Background(), &pb.WatchMovieRequest{
		MovieId: uint32(movieId),
		UserId:  userId.(uint64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
