package internal

import "time"

type Theater struct {
	Id        uint       `json:"id"`
	Name      string     `json:"name"`
	AddressId *uint		 `json:"address_id"`
	Address   *Address   `json:"address"`
	CreatedAt time.Time	 `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at, omitempty"`
	DeletedAt *time.Time `json:"created_at, omitempty"`	
}
