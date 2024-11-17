package movie

import (
	"context"
	"errors"
	"gitlab.com/amarantec/cine/internal"
    "unicode/utf8"
)

type MovieService interface {
	ListMovies(ctx context.Context) ([]internal.Movie, error)
	GetMovieById(ctx context.Context, id uint) (internal.Movie, error)
	AddMovie(ctx context.Context, movie internal.Movie) (uint, error)
	GetMoviesByGenre(ctx context.Context, genre string) ([]internal.Movie, error)
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

func (s *movieService) AddMovie(ctx context.Context, movie internal.Movie) (uint, error) {
    if valid, err := validateMovie(movie); err != nil || !valid {
        return internal.ZERO, err
    }

	return s.movieRepository.AddMovie(ctx, movie)
}

func (s *movieService) GetMoviesByGenre(ctx context.Context, genre string) ([]internal.Movie, error) {
	return s.movieRepository.GetMoviesByGenre(ctx, genre)
}

func validateMovie(m internal.Movie) (bool, error) {
    if m.Title == "" {
        return false, ErrMovieTitleEmpty
    } else if utf8.RuneCountInString(m.Title) < 2 || utf8.RuneCountInString(m.Title) > 50 {
        return false, ErrMovieTitleInvalidFormat
    }
    
    if m.Synopsis == "" {
        return false, ErrMovieSynopsisEmpty
    } else if utf8.RuneCountInString(m.Synopsis) < 2 || utf8.RuneCountInString(m.Title) > 150 {
        return false, ErrMovieSynopsisInvalidFormat
    }

    if len(m.Genre) == internal.ZERO {
        return false, ErrMovieGenreEmpty
    }

    for _, genre := range m.Genre {
        if utf8.RuneCountInString(genre) < 2 || utf8.RuneCountInString(genre) > 20 {
            return false, ErrMovieGenreInvalidFormat
        }
    }

    if len(m.Director) == internal.ZERO {
        return false, ErrMovieDirectorEmpty
    }

    for _, director := range m.Director {
        if utf8.RuneCountInString(director) < 2 || utf8.RuneCountInString(director) > 50 {
            return false, ErrMovieDirectorInvalidFormat
        }
    }

    if len(m.Cast) == internal.ZERO {
        return false, ErrMovieCastEmpty
    }

    for _, cast := range m.Cast {
        if utf8.RuneCountInString(cast) < 2 || utf8.RuneCountInString(cast) > 50 {
            return false, ErrMovieCastInvalidFormat
        }
    }

    if m.ReleaseDate.IsZero() {
        return false, ErrMovieReleaseDateEmpty
    }

    if m.RunningTime == internal.ZERO {
        return false, ErrMovieRunningTimeEmpty
    }

    if m.AgeGroup == internal.ZERO {
        return false, ErrMovieAgeGroupEmpty
    }

    return true, nil
}

var ErrMovieTitleEmpty = errors.New("movie title is empty")
var ErrMovieSynopsisEmpty = errors.New("movie synopsis is empty")
var ErrMovieGenreEmpty = errors.New("movie genre is empty")
var ErrMovieDirectorEmpty = errors.New("movie director is empty")
var ErrMovieCastEmpty = errors.New("movie cast is empty")
var ErrMovieReleaseDateEmpty = errors.New("movie release date is empty")
var ErrMovieRunningTimeEmpty = errors.New("movie running time is empty")
var ErrMovieAgeGroupEmpty = errors.New("movie age group is empty")
var ErrMovieTitleInvalidFormat = errors.New("movie title must be between 2-50 characters")
var ErrMovieSynopsisInvalidFormat = errors.New("movie synopsis must be between 2-150 characters")
var ErrMovieGenreInvalidFormat = errors.New("movie genre must be between 2-20 characters")
var ErrMovieDirectorInvalidFormat = errors.New("movie director must be between 2-50 characters")
var ErrMovieCastInvalidFormat = errors.New("movie cast must be between 2-50 characters")
