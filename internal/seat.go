package internal

import "time"

type Seat struct {
	Row         byte        `json:"row"`
	PlaceNumber uint        `json:"place_number"`
	CineRoom    string      `json:"cine_room"`
    Available   bool        `json:"available"`
    CreatedAt   time.Time   `json:"created_at"`
    UpdatedAt   *time.Time  `json:"updated_at"`
    DeletedAt   *time.Time  `json:"deleted_at"`
}
