package internal

import "time"

type Session struct {
    Id              uint        `json:"id"`
    CineRoomId      uint        `json:"cineroom_id"`
    TheaterId       uint        `json:"theater_id"` 
    MovieId         uint        `json:"movie_id"`
    Schedule        time.Time   `json:"schedule"`
    Date            time.Time   `json:"date"`
    AvailableSeats  uint        `json:"available_seats"`
}
    
