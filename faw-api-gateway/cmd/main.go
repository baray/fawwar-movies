package main

import (
	"log"

	"github.com/fawwar-movies/faw-api-gateway/pkg/config"
	"github.com/fawwar-movies/faw-api-gateway/pkg/movie"
	"github.com/fawwar-movies/faw-api-gateway/pkg/user"
	"github.com/fawwar-movies/faw-api-gateway/pkg/watched"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	//r.MaxMultipartMemory = c.MaxUploadSize.(int) << 20 // [max]MiB

	userService := *user.RegisterRoutes(r, &c)
	movie.RegisterRoutes(r, &c, &userService)
	watched.RegisterRoutes(r, &c, &userService)

	r.Run(c.Port)
}
