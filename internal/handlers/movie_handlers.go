package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"gitlab.com/amarantec/cine/internal/movie"
	"gitlab.com/amarantec/cine/internal"
)

type MovieHandler struct {
	service movie.MovieService
}

func NewMovieHandler(service movie.MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

// listMovies godoc
//
// @Sumary          List all movies
// @Description     List all movies informartion in database
// @Tags            movies
// @Accept          json
// @Produce         json
// @Success         200 {object} internal.Movie
// @Failure         500 {object} string "Err could not list movies"
// @Router          /list-movies [get]
func (h *MovieHandler) listMovies(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	movies, err := h.service.ListMovies(ctxTimeout)
	if err != nil {
		http.Error(w,
			"could not list movies, error: " + err.Error(),
			http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.MarshalIndent(movies, "", "  ")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// getMovieById godoc
//
// @Sumary          Get a movie by Id        
// @Description     Retrieve movie info using its id
// @Tags            movies
// @Accept          json
// @Produce         json
// @Param           id path uint true "Movie Id"
// @Success         200 {object} internal.Movie "Details of movie"
// @Failure         400 {object} string "Err Invalid parameter"
// @Failure         500 {object} string "Err could not get this movie"
// @Router          /get-movie-by-id/{id} [get]
func (h *MovieHandler) getMovieById(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w,
			"invalid parameter, error: " + err.Error(),
			http.StatusBadRequest)
		return
	}

	movie, err := h.service.GetMovieById(ctxTimeout, uint(id))
	if err != nil {
		http.Error(w,
			"could not get this movie, error: " + err.Error(),
			http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.MarshalIndent(movie, "", "  ")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// addMovie godoc
//
// @Sumary          Insert a new movie
// @Description     Insert a new movie in database
// @Tags            movies
// @Accept          json
// @Produce         json
// @Param           movie body internal.Movie true "Movie"
// @Success         201 {object} internal.Movie
// @Failure         400 {object} string "Err could not decote this movie"
// @Failure         500 {object} string "Err when insert the movie"
// @Router          /add-movie [post]
func (h *MovieHandler) addMovie(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	movie := internal.Movie{}

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w,
			"could not decode this request, error: " + err.Error(),
			http.StatusBadRequest)
		return
	}

	response, err := h.service.AddMovie(ctxTimeout, movie)
	if err != nil {
		http.Error(w,
			"could not insert this movie, error: " + err.Error(),
			http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}
		
// getMoviesByGenre godoc
//
// @Sumary          List movies by genre
// @Description     Get a list of movies using its Genre
// @Tags            movies
// @Accept          json
// @Produce         json
// @Param           genre path string true "Movies Genre"
// @Success         200 {object} internal.Movie "Details of movies with this genre"
// @Failure         400 {object} string "Err invalid parameter"
// @Failure         500 {ojbect} string "Err could not list movies"
// @Router          /get-movies-by-genre/{genre} [get]
func (h *MovieHandler) getMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	
	genre := r.PathValue("genre")

	response, err := h.service.GetMoviesByGenre(ctxTimeout, genre)
	if err != nil {
	http.Error(w,
		"could not insert this movie, error: " + err.Error(),
			http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.MarshalIndent(response, "", "  ")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// updateMovie godoc
//
// @Sumary          Update a movie
// @Description     Update a movie info
// @Tags            movies
// @Accept          json
// @Produce         json
// @Param           movie body internal.Movie true "Movie"
// @Success         204 {object} boolean "true"
// @Failure         400 {object} string "Err could not decode this movie"
// @Failure         500 {object} string "Err could not update this movie"
// @Router          /update-movie [put]
func (h *MovieHandler) updateMovie(w http.ResponseWriter, r *http.Request) {
    ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()

    movie := internal.Movie{}

    if err :=
        json.NewDecoder(r.Body).Decode(&movie); err != nil {
            http.Error(w,
                "could not decode this request, error: " + err.Error(),
                http.StatusBadRequest)
            return
    }

    response, err := h.service.UpdateMovie(ctxTimeout, movie)
    if err != nil {
        http.Error(w,
            "could not update this movie, error: " + err.Error(),
            http.StatusInternalServerError)
        return
    }

    jsonResponse, _ := json.Marshal(response)

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusNoContent)
    w.Write(jsonResponse)
}

// deleteMovie godoc
//
// @Sumary          Delete movie
// @Description     Delete a movie register
// @Tags            movies
// @Accept          json
// @Produce         json
// @Param           id path uint true "Movie Id"
// @Success         204 {object} boolean "Movie Id"
// @Failure         400 {object} string "Err invalid parameter"
// @Failure         500 {object} string "Err could not delete this movie"
// @Router          /delete-movie/{movieId} [delete]
func (h *MovieHandler) deleteMovie(w http.ResponseWriter, r *http.Request) {
    ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()
    
    id, err := strconv.Atoi(r.PathValue("movieId"))
    if err != nil {
        http.Error(w,
            "invalid paramter, error: " + err.Error(),
            http.StatusBadRequest)
        return
    }

    response, err := h.service.DeleteMovie(ctxTimeout, uint(id))
    if err != nil {
        http.Error(w,
            "could not delete this movie, error: " + err.Error(),
            http.StatusInternalServerError)
        return
    }

    jsonResponse, _ := json.Marshal(response)

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusNoContent)
    w.Write(jsonResponse)
}
