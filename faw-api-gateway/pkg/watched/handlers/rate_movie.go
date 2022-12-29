package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/fawwar-movies/faw-api-gateway/pkg/watched/pb"
	"github.com/gin-gonic/gin"
)

type RateMovieRequestBody struct {
	Rating uint8  `json:"rating" binding:"required,gte=1,lte=5"`
	Review string `json:"review" mod:"trim"`
}

func RatingMovie(ctx *gin.Context, c pb.WatchedServiceClient) {
	movieId, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	b := RateMovieRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("currentUserId")

	res, err := c.RatingMovie(context.Background(), &pb.RatingMovieRequest{
		MovieId: uint32(movieId),
		Rating:  uint32(b.Rating),
		Review:  b.Review,
		UserId:  userId.(uint64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
