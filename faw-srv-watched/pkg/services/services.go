package services

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"

	"github.com/fawwar-movies/faw-srv-movie/pkg/db"
	"github.com/fawwar-movies/faw-srv-movie/pkg/models"
	"github.com/fawwar-movies/faw-srv-watched/pkg/pb"

	"gorm.io/gorm"
)

type Server struct {
	H db.Handler
	pb.UnimplementedWatchedServiceServer
}

func getStatus(code int) string {
	return http.StatusText(code)
}

func (s *Server) WatchMovie(ctx context.Context, req *pb.WatchMovieRequest) (*pb.StatusResponse, error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{Id: req.UserId}).First(&user); result.Error != nil {
		return &pb.StatusResponse{
			Status: getStatus(http.StatusNotFound),
			Error:  "User not found",
		}, nil
	}

	var movie models.Movie

	if result := s.H.DB.Where(&models.Movie{Id: req.MovieId}).First(&movie); result.Error != nil {
		return &pb.StatusResponse{
			Status: getStatus(http.StatusNotFound),
			Error:  "Movie not found",
		}, nil
	}

	if result := s.H.DB.Model(&user).Association("Watches").Append(&movie); result.Error == nil {
		return &pb.StatusResponse{
			Status: getStatus(http.StatusConflict),
			Error:  "Already watched",
		}, nil
	}

	return &pb.StatusResponse{
		Status: getStatus(http.StatusCreated),
	}, nil
}

func (s *Server) WatchedList(ctx context.Context, req *pb.WatchedListRequest) (*pb.WatchedListResponse, error) {
	offset := (req.Page - 1) * req.Size
	order := req.SortBy + " " + req.Direction
	var data []*pb.Movie

	if result := s.H.DB.Model(&models.Movie{}).
		Select("id", "name", "description", "date", "rate", "cover", "user_id").
		Where("user_id = ?", req.UserId).
		Order(order).
		Offset(int(offset)).
		Limit(int(req.Size)).
		Scan(&data); result.Error != nil {
		return &pb.WatchedListResponse{
			Movies: data,
			Status: getStatus(http.StatusNotFound),
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.WatchedListResponse{
		Movies: data,
		Status: getStatus(http.StatusOK),
	}, nil
}

type Ratings struct {
	MovieId uint32
	Five    uint
	Four    uint
	Three   uint
	Two     uint
	One     uint
}

func (s *Server) fetchMovieRatings(mId uint32) Ratings {
	var ratings = Ratings{
		MovieId: mId,
		Five:    0,
		Four:    0,
		Three:   0,
		Two:     0,
		One:     0,
	}

	var reviews []models.Review
	s.H.DB.Where("movie_id = ?", mId).Find(&reviews)
	for _, review := range reviews {
		switch review.Rate {
		case 5:
			ratings.Five++
		case 4:
			ratings.Four++
		case 3:
			ratings.Three++
		case 2:
			ratings.Two++
		case 1:
			ratings.One++
		}
	}
	fmt.Println(ratings)
	return ratings
}

type UserWatchMovie struct {
	UserID  uint64
	MovieID uint32
}

func (s *Server) RatingMovie(ctx context.Context, req *pb.RatingMovieRequest) (*pb.MovieResponse, error) {
	var user models.User
	if result := s.H.DB.Where(&models.User{Id: req.UserId}).First(&user); result.Error != nil {
		return &pb.MovieResponse{
			Status: getStatus(http.StatusNotFound),
			Error:  "User not found",
		}, nil
	}

	var movie *pb.Movie
	if result := s.H.DB.
		Select("id", "name", "description", "date", "rate", "cover", "user_id").
		Where(&models.Movie{Id: req.MovieId}).First(&movie); result.Error != nil {
		return &pb.MovieResponse{
			Status: getStatus(http.StatusNotFound),
			Error:  "Movie not found",
		}, nil
	}

	var userWatch UserWatchMovie
	if result := s.H.DB.Table("user_watches").Select("*").Where("user_id = ? AND movie_id = ?", req.UserId, req.MovieId).First(&userWatch); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &pb.MovieResponse{
			Status: getStatus(http.StatusBadRequest),
			Error:  "You do not have access to rating this movie",
		}, nil
	}

	reviewInsert := models.Review{MovieID: req.MovieId, UserID: req.UserId, Rate: float32(req.Rating), Comment: req.Review}
	if result := s.H.DB.Select("movie_id", "user_id", "rate", "description").Create(&reviewInsert); result.Error != nil {
		return &pb.MovieResponse{
			Status: getStatus(http.StatusNotFound),
			Error:  "Movie not found",
		}, nil
	}

	var avgRating float64
	ratings := s.fetchMovieRatings(req.MovieId)
	avgRating = math.Round(float64(ratings.Five*5+ratings.Four*4+ratings.Three+ratings.Two*2+ratings.One) / float64(ratings.Five+ratings.Four+ratings.Three+ratings.Two+ratings.One))

	if result := s.H.DB.Model(&models.Movie{}).Where("id = ?", req.MovieId).UpdateColumn("rate", float32(avgRating)); result.Error == nil {
		return &pb.MovieResponse{
			Status: getStatus(http.StatusConflict),
			Error:  "Already watched",
		}, nil
	}

	return &pb.MovieResponse{
		Movie:  movie,
		Status: getStatus(http.StatusCreated),
	}, nil
}
