package internal

import "time"

type Address struct {
	Id        uint   	 `json:"id"`
	City      string 	 `json:"city"`
	Street    string 	 `json:"street"`
	ZIP       string 	 `json:"zip"`
	State     string 	 `json:"state"`
	Country   string 	 `json:"country"`
	CreatedAt time.Time	 `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
