package internal

import "time"

// swagger: model Movie
// @description Movie struct
type Movie struct {
	Id          uint      	`json:"id"`
	Title       string    	`json:"title" example:"Batman: The Dark Knight" minLength:"2" maxLength:"50"`
	Synopsis    string    	`json:"synopsis" example:"The movie begins with a gang of men with clown masks..." minLength:"2" maxLength:"150"`
	Genre       []string  	`json:"genre" example:"'Action', 'Drama'" minLength:"2" maxLength:"20"`
	Director    []string  	`json:"director" example:"Christopher Nolan" minLength:"2" maxLength:"50"`
	Cast        []string  	`json:"cast" example:"'Christian Bale', 'Heath Ledger'" minLength:"2" maxLength:"50"`
	ReleaseDate time.Time 	`json:"release_date" example:"2008-07-14"`
	RunningTime string    	`json:"running_time" example:"120" min:"1" max:"440"`
	AgeGroup    string     	`json:"age_group" example:"12" min:"L" max:"18"`
	CreatedAt	time.Time 	`json:"-"`	
	UpdatedAt	*time.Time  `json:"-"`
	DeletedAt	*time.Time  `json:"-"`	
} // @name Movie
