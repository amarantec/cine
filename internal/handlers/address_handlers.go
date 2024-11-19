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

// AddressHandler is responsible for managing addresses in the API.
type AddressHandler struct {
	service address.AddressService
}

// NewAddressHandler create a new AddressHandler
func NewAddressHandler(service address.AddressService) *AddressHandler {
	return &AddressHandler{service: service}
}

// insertAddress godoc
//
// @Sumary          Insert a new address
// @Description     Insert a new address in database
// @Tags            addresses
// @Accept          json
// @Produce         json
// @Param           address body internal.Address true "Address"
// @Success         201 {object} internal.Address
// @Failure         400 {object} string "Err when decoding the address"
// @Failure         500 {object} string "Err when insert the address"
// @Router          /insert-address [post]
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

// getAddress godoc
//
// @Sumary          Get a address by Id
// @Description     Retrieve address info using its Id
// @Tags            addresses
// @Accept          json
// @Produce         json
// @Param           id path uint true "Address Id"
// @Success         200 {object} internal.Address "Details of the address"
// @Failure         400 {object} string "Invalid parameter"
// @Failure         500 {object} string "Err could not get this address"
// @Router          /get-address/{id} [get]
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

// updateAddress godoc
//
// @Sumary          Update a address 
// @Description     Update address info 
// @Tags            addresses
// @Accept          json
// @Produce         json
// @Param           address body internal.Address true "Address"
// @Success         204 {object} boolean "true"
// @Failure         400 {object} string "Err could not parse this request"
// @Failure         500 {object} string "Err could not update this address"
// @Router          /update-address [put]
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

// deleteAddress godoc
//
// @Sumary          Delete a address 
// @Description     Deletete address register 
// @Tags            addresses
// @Accept          json
// @Produce         json
// @Param           id path uint true "Address Id"
// @Success         204 {object} boolean "true"
// @Failure         400 {object} string "Err invalid parameter"
// @Failure         500 {object} string "Err could not delete this address"
// @Router          /delete-address/{id} [delete]
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

