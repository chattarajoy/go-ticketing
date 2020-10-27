package movie

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"commerceiq.ai/ticketing/pkgs/models"
)

type AddMovieInput struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Duration    time.Duration `json:"duration"`
}

func (am *AddMovieInput) Validate(db *gorm.DB) error {
	if am.Name == "" || am.Description == "" || am.Duration == 0 {
		return fmt.Errorf("name, description and duration are mandatory fields")
	}
	return nil
}

type AddMovieOutput struct {
	Movie models.Movie `json:"movie"`
}

type AddMovieShowInput struct {
	MovieID        int       `json:"movie_id"`
	CinemaScreenID int       `json:"cinema_screen_id"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
}

func (ams *AddMovieShowInput) Validate(db *gorm.DB) error {
	if ams.MovieID == 0 || ams.CinemaScreenID == 0 || ams.StartTime.IsZero() || ams.EndTime.IsZero() {
		return fmt.Errorf("movie_id, cinema_screen_id, end_time, and start_time are mandatory fields")
	}
	return nil
}

type AddMovieShowOutput struct {
	Show models.MovieShow
}

type GetMovieShowInput struct {
	ShowID int `json:"show_id"`
}

func (lms *GetMovieShowInput) Validate(db *gorm.DB) error {
	if lms.ShowID == 0 {
		return fmt.Errorf("show_id is mandatory fields")
	}
	return nil
}

type GetMovieShowOutput struct {
	Show models.MovieShow `json:"show"`
}
