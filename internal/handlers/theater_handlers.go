package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"strconv"
	"gitlab.com/amarantec/cine/internal/theater"
	"gitlab.com/amarantec/cine/internal"
)

type TheaterHandler struct {
	service theater.TheaterService
}

func NewTheaterHandler(service theater.TheaterService) *TheaterHandler {
	return &TheaterHandler{service: service}
}

func (h *TheaterHandler) listTheaters(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	theaters, err := h.service.ListTheaters(ctxTimeout)
	if err != nil {
		http.Error(w,
			"could not list theaters, error: "+err.Error(),
			http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.Marshal(theaters)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *TheaterHandler) getTheaterById(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w,
			"invalid parameter, error: " + err.Error(),
			http.StatusBadRequest)
		return
	}

	theater, err := h.service.GetTheaterById(ctxTimeout, uint(id))
	if err != nil {
		http.Error(w,
			"could not get this theater, error: " + err.Error(),
			http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.MarshalIndent(theater, "", "  ")
	
	w.Header().Set("Content-Type", "application/json; charset=utf-8;")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *TheaterHandler) addTheater(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	theater := internal.Theater{}

	if err := json.NewDecoder(r.Body).Decode(&theater); err != nil {
		http.Error(w,
			"could not parse this request, error: " + err.Error(),
			http.StatusBadRequest)
		return
	}

	response, err := h.service.AddTheater(ctxTimeout, theater)
	if err != nil {
		http.Error(w,
			"could not insert this theater, error: " + err.Error(),
			http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}
