package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/fawwar-movies/faw-api-gateway/pkg/movie/pb"
	"github.com/gin-gonic/gin"
)

func DeleteMovie(ctx *gin.Context, c pb.MovieServiceClient) {
	movieId, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	userId, err := ctx.Get("currentUserId")
	if userId < 1 || err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res, err := c.DeleteMovie(context.Background(), &pb.MovieByUserRequest{
		MovieId: uint32(movieId),
		UserId:  userId.(uint64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
