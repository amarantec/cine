package internal

type CineRoom struct {
	Id        uint   `json:"id"`
	Number    string `json:"number"`
	TheaterId uint   `json:"theater_id"`
	Seats     []Seat `json:"seats"`
}
