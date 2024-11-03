package movie

import (
	"context"

	"gitlab.com/amarantec/cine/internal"
)

type MovieService interface {
	ListMovies(ctx context.Context) ([]internal.Movie, error)
	GetMovieById(ctx context.Context, id uint) (internal.Movie, error)
}

type movieService struct {
	movieRepository MovieRepository
}

func NewMovieService(repository MovieRepository) MovieService {
	return &movieService{movieRepository: repository}
}

func (s *movieService) ListMovies(ctx context.Context) ([]internal.Movie, error) {
	return s.movieRepository.ListMovies(ctx)
}

func (s *movieService) GetMovieById(ctx context.Context, id uint) (internal.Movie, error) {
	return s.movieRepository.GetMovieById(ctx, id)
}
