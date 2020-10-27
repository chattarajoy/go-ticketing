package booking

import (
	"fmt"

	"gorm.io/gorm"

	"commerceiq.ai/ticketing/pkgs/models"
)

type BookSeatsInput struct {
	ShowID      int   `json:"show_id"`
	SeatNumbers []int `json:"seat_numbers"`
	UserID      int   `json:"user_id"`
}

func (bs *BookSeatsInput) Validate(db *gorm.DB) error {
	if bs.ShowID == 0 || bs.UserID == 0 || bs.SeatNumbers == nil {
		return fmt.Errorf("show_id, seat_numbers and user_id are required parameters")
	}

	var show models.MovieShow
	var user models.User
	if result := db.Find(&show, bs.ShowID); result.Error != nil {
		return result.Error
	}
	if result := db.Find(&user, bs.ShowID); result.Error != nil {
		return result.Error
	}
	return nil
}

type BookSeatsOutput struct {
	Booking models.Booking `json:"booking"`
}

type ListBookingsOutput struct {
	Bookings []models.Booking `json:"bookings"`
}
