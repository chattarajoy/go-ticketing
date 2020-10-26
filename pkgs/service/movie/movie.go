package movie

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

func (s *Service) AddMovie(inp *AddMovieInput) (*AddMovieOutput, error) {
	movie := models.Movie{
		Name:        inp.Name,
		Description: inp.Description,
		Duration:    inp.Duration,
	}

	if result := s.db.Create(&movie); result.Error != nil {
		return nil, result.Error
	}
	return &AddMovieOutput{Movie: movie}, nil
}

func (s *Service) AddMovieShow(inp *AddMovieShowInput) (*AddMovieShowOutput, error) {
	show := models.MovieShow{
		StartTime:      inp.StartTime,
		EndTime:        inp.EndTime,
		MovieID:        inp.MovieID,
		CinemaScreenID: inp.CinemaScreenID,
	}

	if result := s.db.Create(&show); result.Error != nil {
		return nil, result.Error
	}
	return &AddMovieShowOutput{Show: show}, nil
}

func (s *Service) ListMovieShow(inp *ListMovieShowInput) (*ListMovieShowOutput, error) {

	var show models.MovieShow
	if result := s.db.Preload("CinemaScreen").Preload("BookingSeat").Find(&show, inp.ShowID); result.Error != nil {
		return nil, result.Error
	}
	return &ListMovieShowOutput{Show: show}, nil
}
