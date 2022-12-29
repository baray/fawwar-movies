package handlers

import (
	"context"
	"net/http"

	"github.com/fawwar-movies/faw-api-gateway/pkg/user/pb"
	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	Fullname string `json:"fullname"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, c pb.UserServiceClient) {
	b := RegisterRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Fullname: b.Fullname,
		Age:      uint32(b.Age),
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
