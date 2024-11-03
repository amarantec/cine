package handlers

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/amarantec/cine/internal/movie"
	"gitlab.com/amarantec/cine/internal/theater"
)

func SetRoutes(conn *pgxpool.Pool) *http.ServeMux {
	m := http.NewServeMux()

	// Movie dependency injection
	movieRepository := movie.NewMovieRepository(conn)
	movieService := movie.MovieService(movieRepository)
	movieHandler := NewMovieHandler(movieService)
	// Theater dependency injection
	theaterRepository := theater.NewTheaterRepository(conn)
	theaterService := theater.NewTheaterService(theaterRepository)
	theaterHandler := NewTheaterHandler(theaterService)

	/*
		ROUTES
	*/

	m.HandleFunc("/movies/list-movies", movieHandler.listMovies)
	m.HandleFunc("/movies/get-movie-by-id/{id}", movieHandler.getMovieById)
	m.HandleFunc("/theaters/list-theaters", theaterHandler.listTheaters)
	return m
}
