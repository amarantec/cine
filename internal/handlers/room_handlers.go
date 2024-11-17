package handlers

import (
	"time"
	"context"
	"gitlab.com/amarantec/cine/internal/room"
	"encoding/json"
	"net/http"
	"strconv"
)

type RoomHandler struct {
	service room.RoomService
}

func NewRoomHandler(service room.RoomService) *RoomHandler {
	return &RoomHandler{service: service}
}

func (h *RoomHandler) listRooms(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	theaterId, err := strconv.Atoi(r.PathValue("theaterId"))	
	if err != nil {
		http.Error(w,
			"invalid parameter, error: " + err.Error(),
			http.StatusBadRequest)
		return
	}

	response, err := h.service.ListRooms(ctxTimeout, uint(theaterId))
	if err != nil {
		http.Error(w,
			"could not list rooms, error: " + err.Error(),
			http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.MarshalIndent(response, "", "  ")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *RoomHandler) getRoomById(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	theaterId, err1 := strconv.Atoi(r.PathValue("theaterId"))	
	roomId, err2 := strconv.Atoi(r.PathValue("roomId"))	

	if err1 != nil || err2 != nil {
		http.Error(w,
			"invalid parameter",
			http.StatusBadRequest)
		return
	}

	response, err := h.service.GetRoomById(ctxTimeout, uint(theaterId), uint(roomId))
	if err != nil {
		http.Error(w, 
			"could not get this room, error: " + err.Error(),
			http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.MarshalIndent(response, "", "  ")
	
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *RoomHandler) listAvailableRoomSeats(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	theaterId, err := strconv.Atoi(r.PathValue("theaterId"))	
	if err != nil {
		http.Error(w,
			"invalid parameter, error: " + err.Error(),
			http.StatusBadRequest)
		return
	}

	roomNumber := r.PathValue("roomNumber")

	response, err := h.service.ListAvailableRoomSeats(ctxTimeout, uint(theaterId), roomNumber)
	if err != nil {
		http.Error(w,
			"could not list available room seats, error: " + err.Error(),
			http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.MarshalIndent(response, "", "  ")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
