package cinema

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

func (s *Service) AddCinema(inp *AddCinemaInput) (*AddCinemaOutput, error) {
	c := models.Cinema{
		Name:   inp.CinemaName,
		CityID: inp.CityID,
	}
	result := s.db.Create(&c)
	if result.Error != nil {
		return nil, result.Error
	}
	// added successfully
	s.cache.Delete("ListCinemasOutput")
	return &AddCinemaOutput{Cinema: c}, nil
}

func (s *Service) ListCinemas() (*ListCinemasOutput, error) {
	if cachedValue, ok := s.cache.Get("ListCinemasOutput"); ok {
		return cachedValue.(*ListCinemasOutput), nil
	}
	var cinemas []models.Cinema

	// Get all records
	result := s.db.Preload("CinemaScreens.CinemaSeats").Preload(clause.Associations).Find(&cinemas)
	if result.Error != nil {
		return nil, result.Error
	}
	out := &ListCinemasOutput{Cinemas: cinemas}
	s.cache.Set("ListCinemasOutput", out, time.Duration(10*time.Minute))
	return out, nil
}

func (s *Service) AddCinemaScreen(inp *AddCinemaScreenInput) (*AddCinemaScreenOutput, error) {
	screen := models.CinemaScreen{
		Name:     inp.ScreenName,
		CinemaID: inp.CinemaID,
	}
	result := s.db.Create(&screen)
	if result.Error != nil {
		return nil, result.Error
	}

	out := &AddCinemaScreenOutput{CinemaScreen: screen}

	for _, seat := range inp.Seats {
		se := models.CinemaSeat{
			SeatNumber:     seat.SeatNumber,
			Type:           seat.SeatType,
			CinemaScreenID: screen.ID,
		}
		result := s.db.Create(&se)
		if result.Error != nil {
			return out, result.Error
		}
	}

	// re-fetch all cinemas
	if _, ok := s.cache.Get("ListCinemasOutput"); ok {
		s.cache.Delete("ListCinemasOutput")

	}
	result = s.db.Preload("CinemaSeats").Preload("Cinema").Find(&screen, screen.ID)
	out.CinemaScreen = screen
	return out, result.Error
}
