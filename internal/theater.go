package internal

type Theater struct {
	Id      uint    `json:"id"`
	Name    string  `json:"name"`
	Session []Movie `json:"session"`
	Address string  `json:"address"`
}
