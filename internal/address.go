package internal

type Address struct {
	Id        uint   `json:"id"`
	TheaterId uint   `json:"theater_id"`
	City      string `json:"city"`
	Street    string `json:"street"`
	ZIP       string `json:"zip"`
	State     string `json:"state"`
	Country   string `json:"country"`
}
