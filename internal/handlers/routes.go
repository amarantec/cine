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
    movieMux := http.NewServeMux()
	movieRepository := movie.NewMovieRepository(conn)
	movieService := movie.MovieService(movieRepository)
	movieHandler := NewMovieHandler(movieService)
	// Theater dependency injection
    theaterMux := http.NewServeMux()
	theaterRepository := theater.NewTheaterRepository(conn)
	theaterService := theater.NewTheaterService(theaterRepository)
	theaterHandler := NewTheaterHandler(theaterService)

	/*
		ROUTES
	*/

	movieMux.HandleFunc("/list-movies", movieHandler.listMovies)
	movieMux.HandleFunc("/get-movie-by-id/{id}", movieHandler.getMovieById)
	movieMux.HandleFunc("/add-movie", movieHandler.addMovie)
	movieMux.HandleFunc("/get-movies-by-genre/{genre}", movieHandler.getMoviesByGenre)


	theaterMux.HandleFunc("/list-theaters", theaterHandler.listTheaters)


    m.Handle("/movies/", http.StripPrefix("/movies", movieMux))
    m.Handle("/theaters/", http.StripPrefix("/theaters", theaterMux))
	return m
}
