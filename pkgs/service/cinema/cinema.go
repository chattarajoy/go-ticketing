package cinema

import (
	"gorm.io/gorm"

	"commerceiq.ai/ticketing/internal/cache"
	"commerceiq.ai/ticketing/pkgs/models"
)

type Service struct {
	db    *gorm.DB
	cache cache.Cache
}

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
	return &AddCinemaOutput{Cinema: c}, nil
}

func (s *Service) ListCinemas() (*ListCinemasOutput, error) {
	var cinemas []models.Cinema

	// Get all records
	result := s.db.Preload("City").Find(&cinemas)
	if result.Error != nil {
		return nil, result.Error
	}
	return &ListCinemasOutput{Cinemas: cinemas}, nil
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
			SeatNumber: seat.SeatNumber,
			Type:       seat.SeatType,
		}
		result := s.db.Create(&se)
		if result.Error != nil {
			return out, result.Error
		}
	}

	// re-fetch all cinemas
	result = s.db.Preload("CinemaSeat").Preload("Cinema").Find(&screen, screen.ID)
	return out, result.Error
}
