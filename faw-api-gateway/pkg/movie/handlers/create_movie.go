package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/fawwar-movies/faw-api-gateway/pkg/movie/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreateMovieRequestBody struct {
	Name        string    `json:"name" binding:"required,min=3,max=255"`
	Description string    `json:"description,omitempty" binding:"max=1024"`
	Date        time.Time `json:"date" binding:"required"`
}

func CreateMovie(ctx *gin.Context, c pb.MovieServiceClient) {
	b := CreateMovieRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, err := ctx.Get("currentUserId")
	if userId < 1 || err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res, err := c.CreateMovie(context.Background(), &pb.CreateMovieRequest{
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
