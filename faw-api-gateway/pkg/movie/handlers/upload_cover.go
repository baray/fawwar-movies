package handlers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/fawwar-movies/faw-api-gateway/pkg/movie/pb"
	"github.com/gin-gonic/gin"
)

type UploadCoverRequestBody struct {
	Cover *multipart.FileHeader `form:"cover" binding:"required"`
}

func UploadCover(ctx *gin.Context, c pb.MovieServiceClient) {
	movieId, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	b := UploadCoverRequestBody{}
	if err := ctx.ShouldBind(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if matched, _ := regexp.MatchString("image/j", b.Cover.Header.Values("content-type")[0]); !matched {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"content-type": b.Cover.Header.Values("content-type")[0],
			"error":        "file not image or not supported",
		})
		return
	}

	fileExt := filepath.Ext(b.Cover.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(b.Cover.Filename), filepath.Ext(b.Cover.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt

	if err := ctx.SaveUploadedFile(b.Cover, "./media/"+filename); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	userId, _ := ctx.Get("currentUserId")

	res, err := c.UpdateCover(ctx, &pb.UpdateCoverRequest{
		MovieId:   uint32(movieId),
		CoverName: filename,
		UserId:    userId.(uint64),
	})

	if err != nil {
		// todo: remove image
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
