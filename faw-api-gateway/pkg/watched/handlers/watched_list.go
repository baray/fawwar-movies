package handlers

import (
	"context"
	"net/http"
	"strings"

	movieHandler "github.com/fawwar-movies/faw-api-gateway/pkg/movie/handlers"
	"github.com/fawwar-movies/faw-api-gateway/pkg/pagination"
	"github.com/fawwar-movies/faw-api-gateway/pkg/util"
	"github.com/fawwar-movies/faw-api-gateway/pkg/watched/pb"
	"github.com/gin-gonic/gin"
)

func WatchedList(ctx *gin.Context, c pb.WatchedServiceClient) {
	sortBy := strings.ToLower(ctx.DefaultQuery("sortBy", movieHandler.DefaultSortBy))
	direction := strings.ToLower(ctx.DefaultQuery("direction", movieHandler.DefaultDir))
	pageNum, pageSize := pagination.GetPaginationParametersFromRequest(ctx)

	if util.Contains(movieHandler.SortAvailable, sortBy) == false ||
		util.Contains(movieHandler.DirAvailable, direction) == false {
		ctx.AbortWithStatus(http.StatusBadGateway)
		return
	}

	userId, _ := ctx.Get("currentUserId")

	res, err := c.WatchedList(context.Background(), &pb.WatchedListRequest{
		SortBy:    sortBy,
		Direction: direction,
		Page:      uint32(pageNum),
		Size:      uint32(pageSize),
		UserId:    userId.(uint64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	response := pagination.New(pageNum, pageSize)
	response.Items = res

	ctx.JSON(http.StatusOK, &response)
}
