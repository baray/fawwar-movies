package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ModelDates struct {
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

type Movie struct {
	ModelDates
	Id          uint32    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"column:name; not null; index:uidx_movies_name,unique; size:100"`
	Description string    `json:"description" gorm:"column:description; size:500"`
	Date        time.Time `json:"date" gorm:"column:date; type:timestamp; not null; index:idx_movie_filter,priority:2"`
	Rate        float32   `json:"rate" gorm:"column:rate; type:float; not null; default: 0.0; index:idx_movie_filter,priority:1"`
	Cover       string    `json:"cover" gorm:"column:cover; size:200"`
	UserId      uint64    `json:"userId" gorm:"column:user_id"`
	Users       []*User   `gorm:"many2many:user_watches;"`
	Reviews     []Review  `json:"-"`
}

func CreateOrMigrateMovie(db *gorm.DB) {
	if db.Migrator().HasTable(&Movie{}) {
		fmt.Println("Movie table already exist!")
		db.Migrator().AutoMigrate(&Movie{})
	} else {
		fmt.Println("The data does not exist, so I need to create the movie table first")
		err := db.Migrator().CreateTable(&Movie{}).Error
		if err == nil {
			fmt.Println("Movie table already exist!")
		}
	}
}

type User struct {
	ModelDates
	Id       uint64   `json:"id" gorm:"primaryKey"`
	Fullname string   `json:"fullname" gorm:"column:name; not null; size:100"`
	Email    string   `json:"email" gorm:"column:email; index:uidx_user_email,unique; size:100"`
	Age      uint8    `json:"age" gorm:"column:age"`
	Password string   `json:"password" gorm:"column:password; not null"`
	Movies   []Movie  `json:"-"`
	Watches  []*Movie `json:"-" gorm:"many2many:user_watches;"`
	Reviews  []Review `json:"-"`
}

func CreateOrMigrateUser(db *gorm.DB) {
	if db.Migrator().HasTable(&User{}) {
		fmt.Println("User table already exist!")
		db.Migrator().AutoMigrate(&User{})
	} else {
		fmt.Println("The data does not exist, so I need to create the user table first")
		err := db.Migrator().CreateTable(&User{}).Error
		if err == nil {
			fmt.Println("User table already exist!")
		}
	}
}

type Review struct {
	ModelDates
	Id      uint64  `json:"id" gorm:"primaryKey"`
	Rate    float32 `json:"rate" gorm:"column:rate; not null; default: 0.0; type:float; index:idx_movie_rate"`
	Comment string  `json:"description" gorm:"column:description; size:500"`
	UserID  uint64  `json:"userId" gorm:"column:user_id"`
	MovieID uint32  `json:"movieId" gorm:"column:movie_id"`
}

func CreateOrMigrateReview(db *gorm.DB) {
	if db.Migrator().HasTable(&Review{}) {
		fmt.Println("Review table already exist!")
		db.Migrator().AutoMigrate(&Review{})
	} else {
		fmt.Println("The data does not exist, so I need to create the review table first")
		err := db.Migrator().CreateTable(&Review{}).Error
		if err == nil {
			fmt.Println("Review table already exist!")
		}
	}
}
