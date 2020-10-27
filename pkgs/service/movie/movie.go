package movie

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/chattarajoy/go-ticketing/internal/cache"
	"github.com/chattarajoy/go-ticketing/pkgs/models"
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
	out := &AddMovieShowOutput{Show: show}
	if result := s.db.Preload("Movie").Preload("Bookings").Preload("Seats").First(&show, show.ID); result.Error != nil {
		return nil, result.Error
	}
	out.Show = show
	return out, nil
}

func (s *Service) GetMovieShow(inp *GetMovieShowInput) (*GetMovieShowOutput, error) {

	var show models.MovieShow
	if result := s.db.Preload(clause.Associations).Find(&show, inp.ShowID); result.Error != nil {
		return nil, result.Error
	}
	return &GetMovieShowOutput{Show: show}, nil
}
