package internal

import "time"

type Ticket struct {
	Id              uint        `json:"ticket"`
    SessionId       uint        `json:"session_id"`
    UserId          uint        `json:"user_id"`
	Seat            Seat        `json:"seat"`
	PurchaseDate    time.Time   `json:"purchase_date"`
}
