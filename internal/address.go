package internal

import "time"

// swagger: model Address
// @description Address struct
type Address struct {
	Id        uint   	 `json:"id"`
	City      string 	 `json:"city" example:"Osório" minLength:"2" maxLength:"50"`
	Street    string 	 `json:"street" example:"Costa Gama" minLength:"2" maxLength:"50"`
	ZIP       string 	 `json:"zip" example:"95520000" Length:"8"`
	State     string 	 `json:"state" example:"RS" Length:"2"`
	Country   string 	 `json:"country" example:"BR" Length:"2"`
	CreatedAt time.Time	 `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
} // @name Address
