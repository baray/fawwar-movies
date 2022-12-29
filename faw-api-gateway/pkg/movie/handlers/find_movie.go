package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/fawwar-movies/faw-api-gateway/pkg/movie/pb"
)

func FindMovie(ctx *gin.Context, c pb.MovieServiceClient) {
	movieId, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.FindMovie(context.Background(), &pb.FindMovieRequest{
		Id: uint32(movieId),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
