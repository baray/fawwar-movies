package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/fawwar-movies/faw-api-gateway/pkg/movie/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UpdateMovieRequestBody struct {
	Name        string    `json:"name" binding:"required,min=3,max=255"`
	Description string    `json:"description,omitempty" binding:"max=1024"`
	Date        time.Time `json:"date" binding:"required"`
}

func UpdateMovie(ctx *gin.Context, c pb.MovieServiceClient) {
	movieId, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	b := UpdateMovieRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("currentUserId")

	res, err := c.UpdateMovie(context.Background(), &pb.UpdateMovieRequest{
		MovieId:     uint32(movieId),
		Name:        b.Name,
		Description: b.Description,
		Date:        timestamppb.New(b.Date),
		UserId:      userId.(uint64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
