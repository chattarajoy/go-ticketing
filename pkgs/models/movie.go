package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Movie: the table that stores all the movies
type Movie struct {
	Model

	Name        string        `json:"name"`
	Description string        `json:"description"`
	Duration    time.Duration `json:"duration"`

	// relations
	MovieShows []MovieShow `json:"shows" gorm:"foreignKey:MovieID"`
}

// MovieShows: each screening of a movie
type MovieShow struct {
	Model
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`

	// relations
	CinemaScreenID int `json:"-"`
	MovieID        int `json:"-"`
	CinemaScreen   CinemaScreen
	Movie          Movie
	Bookings       []Booking     `json:"bookings" gorm:"foreignKey:MovieShowID"`
	Seats          []BookingSeat `json:"seats" gorm:"foreignKey:MovieShowID"`
}

// check if show is valid
func (ms *MovieShow) BeforeCreate(db *gorm.DB) (err error) {
	if err := ms.CheckOverlap(db); err != nil {
		return err
	}
	return nil
}

func (ms *MovieShow) CheckOverlap(db *gorm.DB) error {
	var shows []*MovieShow

	if result := db.Where("movie_id = ? AND cinema_screen_id = ? AND ((start_time BETWEEN ? AND ? ) OR (end_time BETWEEN ? AND ? ))",
		ms.MovieID, ms.CinemaScreenID, ms.StartTime, ms.EndTime, ms.StartTime, ms.EndTime).Find(&shows); result.Error != nil {
		return result.Error
	} else {
		if len(shows) > 0 {
			return fmt.Errorf("the show has an overlap with %d shows", len(shows))
		}
	}
	return nil
}

// generate seats
func (ms *MovieShow) AfterCreate(db *gorm.DB) (err error) {
	if err := ms.GenerateShowSeats(db); err != nil {
		return err
	}
	return nil
}
func (ms *MovieShow) GenerateShowSeats(db *gorm.DB) error {
	var seats []*CinemaSeat
	result := db.Where("cinema_screen_id = ?", ms.CinemaScreenID).Find(&seats)

	var showSeats []*BookingSeat
	for _, seat := range seats {
		showSeats = append(showSeats, &BookingSeat{
			Status:       SeatAvailable,
			MovieShowID:  ms.ID,
			CinemaSeatID: seat.ID,
		})
	}
	if result = db.Create(showSeats); result.Error != nil {
		return result.Error
	}
	return nil
}
