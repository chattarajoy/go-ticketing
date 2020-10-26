package booking

import (
	"gorm.io/gorm"

	"commerceiq.ai/ticketing/internal/cache"
	"commerceiq.ai/ticketing/pkgs/models"
)

type Service struct {
	db    *gorm.DB
	cache cache.Cache
}

// TODO: Add Logger
func NewService(db *gorm.DB, cache cache.Cache) *Service {
	return &Service{
		db:    db,
		cache: cache,
	}
}

func (s *Service) BookSeats(inp *BookSeatsInput) (*BookSeatsOutput, error) {
	booking := models.Booking{
		SeatCount:   len(inp.SeatNumbers),
		Status:      models.BookingPending,
		UserID:      inp.UserID,
		MovieShowID: inp.ShowID,
	}
	if result := s.db.Create(&booking); result.Error != nil {
		return nil, result.Error
	}
	out := &BookSeatsOutput{Booking: booking}
	err := booking.BookSeats(s.db, inp.SeatNumbers)
	if err != nil {
		booking.Fail()
	} else {
		booking.Confirm()
	}
	if result := s.db.Save(&booking); result.Error != nil {
		return out, result.Error
	}
	if result := s.db.Preload("BookingSeat").
		Preload("MovieShow").
		Preload("User").
		Find(&booking, booking.ID); result.Error != nil {
		return out, result.Error
	}

	out.Booking = booking
	return out, nil
}
