package internal

type Theater struct {
	Id        uint       `json:"id"`
	Name      string     `json:"name"`
	Address   Address    `json:"address"`
}
