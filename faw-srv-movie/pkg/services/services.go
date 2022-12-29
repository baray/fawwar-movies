package services

import (
	"context"
	"net/http"
	"time"

	"github.com/fawwar-movies/faw-srv-movie/pkg/db"
	"github.com/fawwar-movies/faw-srv-movie/pkg/models"
	pb "github.com/fawwar-movies/faw-srv-movie/pkg/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	H db.Handler
	pb.UnimplementedMovieServiceServer
}

func getStatus(code int) string {
	return http.StatusText(code)
}

func (s *Server) CreateMovie(ctx context.Context, req *pb.CreateMovieRequest) (*pb.StatusResponse, error) {
	var movie models.Movie

	movie.Name = req.Name
	movie.Description = req.Description
	movie.Date = req.Date.AsTime()
	movie.UserId = uint64(req.UserId)

	if result := s.H.DB.Create(&movie); result.Error != nil {
		return &pb.StatusResponse{
			Status: getStatus(http.StatusConflict),
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.StatusResponse{
		Status: getStatus(http.StatusCreated),
		// Id:     movie.Id,
	}, nil
}

func (s *Server) DeleteMovie(ctx context.Context, req *pb.MovieByUserRequest) (*pb.StatusResponse, error) {
	if result := s.H.DB.Where("user_id = ? AND id = ?", req.UserId, req.MovieId).Delete(&models.Movie{}); result.Error != nil {
		return &pb.StatusResponse{
			Status: getStatus(http.StatusConflict),
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.StatusResponse{
		Status: getStatus(http.StatusOK),
		// Id:     req.MovieId,
	}, nil
}

func (s *Server) FindMovie(ctx context.Context, req *pb.FindMovieRequest) (*pb.MovieResponse, error) {
	var movie models.Movie

	if result := s.H.DB.First(&movie, req.Id); result.Error != nil {
		return &pb.MovieResponse{
			Status: getStatus(http.StatusNotFound),
			Error:  result.Error.Error(),
		}, nil
	}

	data := &pb.Movie{
		Id:          movie.Id,
		Name:        movie.Name,
		Description: movie.Description,
		Date:        timestamppb.New(movie.Date),
		Rate:        movie.Rate,
		Cover:       movie.Cover,
		UserId:      movie.UserId,
	}

	return &pb.MovieResponse{
		Status: getStatus(http.StatusOK),
		Movie:  data,
	}, nil
}

func (s *Server) MoviesList(ctx context.Context, req *pb.MoviesListRequest) (*pb.MoviesListResponse, error) {
	offset := (req.Page - 1) * req.Size
	order := req.SortBy + " " + req.Direction
	var data []*pb.Movie

	if result := s.H.DB.Model(&models.Movie{}).
		Select("id", "name", "description", "date", "rate", "cover", "user_id").
		Order(order).
		Offset(int(offset)).
		Limit(int(req.Size)).
		Scan(&data); result.Error != nil {
		return &pb.MoviesListResponse{
			Movies: data,
			Status: getStatus(http.StatusNotFound),
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.MoviesListResponse{
		Movies: data,
		Status: getStatus(http.StatusOK),
	}, nil
}

type UpdateMovie struct {
	Name        string
	Description string
	Date        time.Time
}

func (s *Server) UpdateMovie(ctx context.Context, req *pb.UpdateMovieRequest) (*pb.StatusResponse, error) {
	movie := &UpdateMovie{
		Name:        req.Name,
		Description: req.Description,
		Date:        req.Date.AsTime(),
	}

	if result := s.H.DB.Model(&models.Movie{}).Where("user_id = ? AND id = ?", req.UserId, req.MovieId).
		Updates(movie); result.Error != nil {
		return &pb.StatusResponse{
			Status: getStatus(http.StatusNotFound),
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.StatusResponse{
		Status: getStatus(http.StatusOK),
	}, nil
}

func (s *Server) UpdateCover(ctx context.Context, req *pb.UpdateCoverRequest) (*pb.StatusResponse, error) {
	if result := s.H.DB.Model(&models.Movie{}).Where("user_id = ? AND id = ?", req.UserId, req.MovieId).
		Updates(map[string]interface{}{"cover": req.CoverName}); result.Error != nil {
		return &pb.StatusResponse{
			Status: getStatus(http.StatusNotFound),
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.StatusResponse{
		Status: getStatus(http.StatusOK),
	}, nil
}
