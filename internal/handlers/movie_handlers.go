package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"gitlab.com/amarantec/cine/internal/movie"
)

type MovieHandler struct {
	service movie.MovieService
}

func NewMovieHandler(service movie.MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

func (h *MovieHandler) listMovies(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	movies, err := h.service.ListMovies(ctxTimeout)
	if err != nil {
		http.Error(w,
			"could not list movies, error: "+err.Error(),
			http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.Marshal(movies)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *MovieHandler) getMovieById(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := strconv.Atoi(r.URL.Path[len("/movies/get-movie-by-id/"):])
	if err != nil {
		http.Error(w,
			"invalid parameter",
			http.StatusBadRequest)
		return
	}

	movie, err := h.service.GetMovieById(ctxTimeout, uint(id))
	if err != nil {
		http.Error(w,
			"could not get this movie",
			http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.Marshal(movie)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
