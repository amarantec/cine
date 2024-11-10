package handlers

import (
	"context"
	"net/http"
	"time"
	"encoding/json"
	"gitlab.com/amarantec/cine/internal"
	"gitlab.com/amarantec/cine/internal/address"
	"strconv"
)

type AddressHandler struct {
	service address.AddressService
}

func NewAddressHandler(service address.AddressService) *AddressHandler {
	return &AddressHandler{service: service}
}

func (h *AddressHandler) insertAddress(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	address := internal.Address{}

	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		http.Error(w,
			"could not decode this request, error: " + err.Error(),
			http.StatusBadRequest)
		return
	}

	response, err := h.service.InsertAddress(ctxTimeout, address)
	if err != nil {
		http.Error(w,
			"could not insert this address, error: " + err.Error(),
			http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")	
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func (h *AddressHandler) getAddress(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w,
			"invalid parameter, error: " + err.Error(),
			http.StatusBadRequest)
		return
	}

	address, err := h.service.GetAddress(ctxTimeout, uint(id))
	if err != nil {
		http.Error(w,
			"could not get this address, error: " + err.Error(),
			http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.MarshalIndent(address, "", "  ")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *AddressHandler) updateAddress(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	address := internal.Address{}

	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		http.Error(w,
			"could not parse this request, error: " + err.Error(),
			http.StatusBadRequest)
		return
	}

	response, err := h.service.UpdateAddress(ctxTimeout, address)
	if err != nil {
		http.Error(w,
			"could not update this address, error:" + err.Error(),
			http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNoContent)
	w.Write(jsonResponse)
}

func (h *AddressHandler) deleteAddress(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w,
			"invalid parameter, error: " + err.Error(),
			http.StatusBadRequest)
		return
	}

	response, err := h.service.DeleteAddress(ctxTimeout, uint(id))
	if err != nil {
		http.Error(w,
			"could not delete this address, error: " + err.Error(),
			http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNoContent)
	w.Write(jsonResponse)
}

