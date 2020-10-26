package cinema

import (
	"fmt"

	"gorm.io/gorm"

	"commerceiq.ai/ticketing/pkgs/models"
)

type AddCinemaInput struct {
	CinemaName string `json:"cinema_name"`
	CityID     int    `json:"city_id"`
}

func (ac *AddCinemaInput) Validate(db *gorm.DB) error {
	// check non-empty valid inputs
	if ac.CinemaName == "" || ac.CityID == 0 {
		return fmt.Errorf("cinema_name and city_id are required parameters")
	}

	// check if city id exists
	var city models.City
	if result := db.First(&city, ac.CityID); result.Error != nil {
		return fmt.Errorf("error fetching city: %s ", result.Error.Error())
	}
	return nil
}

type AddCinemaOutput struct {
	Cinema models.Cinema `json:"cinema"`
}

type ListCinemasOutput struct {
	Cinemas []models.Cinema `json:"cinemas"`
}

type AddCinemaScreenInput struct {
	CinemaID   int         `json:"cinema_id"`
	ScreenName string      `json:"screen_name"`
	Seats      []*SeatInfo `json:"seats"`
}

func (acs *AddCinemaScreenInput) Validate(db *gorm.DB) error {

	// check non-empty valid inputs
	if acs.CinemaID == 0 || acs.ScreenName == "" || acs.Seats == nil {
		return fmt.Errorf("cinema_id , seats and screen_name are required parameters")
	}

	// check if city id exists
	var cinema models.Cinema
	if result := db.First(&cinema, acs.CinemaID); result.Error != nil {
		return result.Error
	}
	return nil
}

type SeatInfo struct {
	SeatNumber int             `json:"seat_number"`
	SeatType   models.SeatType `json:"seat_type"`
}

type AddCinemaScreenOutput struct {
	CinemaScreen models.CinemaScreen `json:"cinema_screen"`
}
