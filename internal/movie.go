package internal

import "time"

type Movie struct {
	Id          uint      `json:"id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Synopsis    string    `json:"synopsis" binding:"required"`
	Genre       []string  `json:"genre" binding:"required"`
	Director    []string  `json:"director" binding:"required"`
	Cast        []string  `json:"cast" binding:"required"`
	ReleaseDate time.Time `json:"release_date" binding:"required"`
	RunningTime uint      `json:"running_time" binding:"required"`
	AgeGroup    uint      `json:"age_group" binding:"required"`
}
