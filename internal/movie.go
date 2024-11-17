package internal

import "time"

type Movie struct {
	Id          uint      	`json:"id"`
	Title       string    	`json:"title"`
	Synopsis    string    	`json:"synopsis"`
	Genre       []string  	`json:"genre"`
	Director    []string  	`json:"director"`
	Cast        []string  	`json:"cast"`
	ReleaseDate time.Time 	`json:"release_date"`
	RunningTime uint      	`json:"running_time"`
	AgeGroup    uint      	`json:"age_group"`
	CreatedAt	time.Time 	`json:"created_at"`	
	UpdatedAt	*time.Time  `json:"updated_at"`
	DeletedAt	*time.Time  `json:"deleted_at"`	
}
