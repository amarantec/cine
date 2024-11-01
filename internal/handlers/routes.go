package handlers

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/amarantec/cine/internal/movie"
)

func SetRoutes(conn *pgxpool.Pool) *http.ServeMux {
	m := http.NewServeMux()

	var movieRepository = movie.NewMovieRepository(conn)
	var movieService = movie.MovieService(movieRepository)
	var movieHandler = NewMovieHandler(movieService)

	m.HandleFunc("/movies/list-movies", movieHandler.ListMovies)

	return m
}
