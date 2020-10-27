package movie

import (
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/chattarajoy/go-ticketing/internal/testhelpers"
)

func TestAddMovieInput_Validate(t *testing.T) {
	db := testhelpers.SetupDB()
	type fields struct {
		Name        string
		Description string
		Duration    time.Duration
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			fields: fields{
				Name:        "test",
				Description: "test",
				Duration:    5000,
			},
			args:    args{db: db},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			am := &AddMovieInput{
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				Duration:    tt.fields.Duration,
			}
			if err := am.Validate(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddMovieShowInput_Validate(t *testing.T) {
	db := testhelpers.SetupDB()
	type fields struct {
		MovieID        int
		CinemaScreenID int
		StartTime      time.Time
		EndTime        time.Time
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			fields: fields{
				MovieID:        1,
				CinemaScreenID: 1,
				StartTime:      time.Now(),
				EndTime:        time.Now().Add(2 * time.Hour),
			},
			args:    args{db: db},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ams := &AddMovieShowInput{
				MovieID:        tt.fields.MovieID,
				CinemaScreenID: tt.fields.CinemaScreenID,
				StartTime:      tt.fields.StartTime,
				EndTime:        tt.fields.EndTime,
			}
			if err := ams.Validate(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetMovieShowInput_Validate(t *testing.T) {
	db := testhelpers.SetupDB()
	type fields struct {
		ShowID int
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Basic",
			fields:  fields{ShowID: 1},
			args:    args{db: db},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lms := &GetMovieShowInput{
				ShowID: tt.fields.ShowID,
			}
			if err := lms.Validate(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
