package cinema

import (
	"testing"

	"gorm.io/gorm"

	"github.com/chattarajoy/go-ticketing/internal/testhelpers"
)

func TestAddCinemaInput_Validate(t *testing.T) {
	db := testhelpers.SetupDB()
	type fields struct {
		CinemaName string
		CityID     int
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
				CinemaName: "test",
				CityID:     1,
			},
			args:    args{db: db},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AddCinemaInput{
				CinemaName: tt.fields.CinemaName,
				CityID:     tt.fields.CityID,
			}
			if err := ac.Validate(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddCinemaScreenInput_Validate(t *testing.T) {
	db := testhelpers.SetupDB()
	type fields struct {
		CinemaID   int
		ScreenName string
		Seats      []*SeatInfo
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
				CinemaID:   1,
				ScreenName: "test",
				Seats:      []*SeatInfo{},
			},
			args:    args{db: db},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acs := &AddCinemaScreenInput{
				CinemaID:   tt.fields.CinemaID,
				ScreenName: tt.fields.ScreenName,
				Seats:      tt.fields.Seats,
			}
			if err := acs.Validate(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
