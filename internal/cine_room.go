package internal

import "time"

type CineRoom struct {
	Id        uint       `json:"id"`
	Number    string     `json:"number"`
	TheaterId uint       `json:"theater_id"`
	Seats     []Seat     `json:"seats"`
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at"`
}
