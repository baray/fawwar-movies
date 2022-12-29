package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/fawwar-movies/faw-api-gateway/pkg/movie/pb"
	"github.com/fawwar-movies/faw-api-gateway/pkg/pagination"
	"github.com/fawwar-movies/faw-api-gateway/pkg/util"
	"github.com/gin-gonic/gin"
)

var (
	SortAvailable = []string{"rate", "date"}
	DirAvailable  = []string{"asc", "desc"}

	DefaultSortBy = "rate"
	DefaultDir    = "desc"
)

func MoviesList(ctx *gin.Context, c pb.MovieServiceClient) {
	sortBy := strings.ToLower(ctx.DefaultQuery("sortBy", DefaultSortBy))
	direction := strings.ToLower(ctx.DefaultQuery("direction", DefaultDir))
	pageNum, pageSize := pagination.GetPaginationParametersFromRequest(ctx)

	if util.Contains(SortAvailable, sortBy) == false ||
		util.Contains(DirAvailable, direction) == false {
		ctx.AbortWithStatus(http.StatusBadGateway)
		return
	}

	res, err := c.MoviesList(context.Background(), &pb.MoviesListRequest{
		SortBy:    sortBy,
		Direction: direction,
		Page:      uint32(pageNum),
		Size:      uint32(pageSize),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	response := pagination.New(pageNum, pageSize)
	response.Items = res

	ctx.JSON(http.StatusOK, &response)
}
