package db

import (
	"log"

	"github.com/fawwar-movies/faw-srv-movie/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(sqlite.Open(url), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	models.CreateOrMigrateMovie(db)
	models.CreateOrMigrateUser(db)
	models.CreateOrMigrateReview(db)

	return Handler{db}
}
