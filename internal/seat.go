package internal

type Seat struct {
	Row         byte   `json:"row"`
	PlaceNumber uint   `json:"place_number"`
	CineRoom    string `json:"cine_room"`
}
