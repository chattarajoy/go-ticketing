package models

import (
	"time"
)

// Movie: the table that stores all the movies
type Movie struct {
	Model

	Name        string
	Description string
	Duration    string

	// relations
	MovieShows []MovieShow `gorm:"foreignKey:MovieID"`
}

// MovieShows: each screening of a movie
type MovieShow struct {
	Model
	ShowDate  time.Time
	StartTime time.Time
	EndTime   time.Time

	// relations
	CinemaScreenID int `json:"cinema_screen_id"`
	MovieID        int `json:"movie_id"`
	CinemaScreen   CinemaScreen
	Movie          Movie
	Bookings       []Booking `gorm:"foreignKey:MovieShowID"`
}
