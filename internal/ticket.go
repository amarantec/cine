package internal

type Ticket struct {
	Id       uint   `json:"ticket"`
	MovieId  uint   `json:"movie_id"`
	Movie    string `json:"movie"`
	CineRoom string `json:"cine_room"`
}
