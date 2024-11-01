package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gitlab.com/amarantec/cine/internal/movie"
)

type MovieHandler struct {
	service movie.MovieService
}

func NewMovieHandler(service movie.MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

func (h *MovieHandler) ListMovies(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	movies, err := h.service.ListMovies(ctxTimeout)
	if err != nil {
		log.Printf("error: %v", err)
		http.Error(w,
			"could not list movies",
			http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.Marshal(movies)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
