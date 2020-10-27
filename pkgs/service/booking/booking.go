package booking

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

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

func (s *Service) ListBookings() (*ListBookingsOutput, error) {
	if cachedValue, ok := s.cache.Get("ListBookingsOutput"); ok {
		return cachedValue.(*ListBookingsOutput), nil
	}
	var bookings []models.Booking

	// Get all records
	result := s.db.Preload(clause.Associations).Find(&bookings)
	if result.Error != nil {
		return nil, result.Error
	}
	out := &ListBookingsOutput{Bookings: bookings}
	s.cache.Set("ListBookingsOutput", out, time.Duration(10*time.Minute))
	return out, nil
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

	// clear cache
	if _, ok := s.cache.Get("ListBookingsOutput"); ok {
		s.cache.Delete("ListBookingsOutput")
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
