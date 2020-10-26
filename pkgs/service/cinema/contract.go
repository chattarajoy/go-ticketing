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
	var city *models.City
	if result := db.Find(&city, ac.CityID); result.Error != nil {
		return result.Error
	}
	return nil
}

type AddCinemaOutput struct {
	Cinema models.Cinema `json:"cinema"`
}

type ListCinemasOutput struct {
	Cinemas []models.Cinema `json:"cinemas"`
}

type AddCinemaScreensInput struct {
	CinemaID int `json:"cinema_id"`
}
