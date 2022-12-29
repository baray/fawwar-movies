package user

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/fawwar-movies/faw-api-gateway/pkg/user/pb"
)

type UserMiddlewareConfig struct {
	srv *ServiceClient
}

func InitUserMiddleware(srv *ServiceClient) UserMiddlewareConfig {
	return UserMiddlewareConfig{srv}
}

func (c *UserMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.srv.Client.Authenticated(context.Background(), &pb.AuthRequest{
		Token: strings.TrimSpace(token[1]),
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("currentUserId", res.UserId)

	ctx.Next()
}
