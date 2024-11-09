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
