package cinema

import (
	"reflect"
	"testing"

	"gorm.io/gorm"

	"github.com/chattarajoy/go-ticketing/internal/cache"
	"github.com/chattarajoy/go-ticketing/internal/testhelpers"
	"github.com/chattarajoy/go-ticketing/pkgs/models"
)

func TestNewService(t *testing.T) {
	db := testhelpers.SetupDB()
	cacheOb := cache.NewCache(cache.InMemoryCache)

	type args struct {
		db    *gorm.DB
		cache cache.Cache
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		{
			name: "TestNew",
			args: args{
				db:    db,
				cache: cacheOb,
			},
			want: NewService(db, cacheOb),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.db, tt.args.cache); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_AddCinema(t *testing.T) {
	db := testhelpers.SetupDB()
	cacheOb := cache.NewCache(cache.InMemoryCache)

	type fields struct {
		db    *gorm.DB
		cache cache.Cache
	}
	type args struct {
		inp *AddCinemaInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AddCinemaOutput
		wantErr bool
	}{
		{
			name: "Add",
			fields: fields{
				db:    db,
				cache: cacheOb,
			},
			args: args{inp: &AddCinemaInput{
				CinemaName: "test",
				CityID:     2,
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				db:    tt.fields.db,
				cache: tt.fields.cache,
			}
			got, err := s.AddCinema(tt.args.inp)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddCinema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddCinema() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_AddCinemaScreen(t *testing.T) {
	db := testhelpers.SetupDB()
	cacheOb := cache.NewCache(cache.InMemoryCache)

	type fields struct {
		db    *gorm.DB
		cache cache.Cache
	}
	type args struct {
		inp *AddCinemaScreenInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AddCinemaScreenOutput
		wantErr bool
	}{
		{
			name: "Test",
			fields: fields{
				db:    db,
				cache: cacheOb,
			},
			args: args{inp: &AddCinemaScreenInput{
				CinemaID:   1,
				ScreenName: "test",
				Seats:      nil,
			}},
			wantErr: true,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				db:    tt.fields.db,
				cache: tt.fields.cache,
			}
			got, err := s.AddCinemaScreen(tt.args.inp)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddCinemaScreen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddCinemaScreen() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ListCinemas(t *testing.T) {
	db := testhelpers.SetupDB()
	cacheOb := cache.NewCache(cache.InMemoryCache)

	type fields struct {
		db    *gorm.DB
		cache cache.Cache
	}
	tests := []struct {
		name    string
		fields  fields
		want    *ListCinemasOutput
		wantErr bool
	}{
		{
			name: "Basic",
			fields: fields{
				db:    db,
				cache: cacheOb,
			},
			want:    &ListCinemasOutput{Cinemas: []models.Cinema{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				db:    tt.fields.db,
				cache: tt.fields.cache,
			}
			got, err := s.ListCinemas()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListCinemas() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListCinemas() got = %v, want %v", got, tt.want)
			}
		})
	}
}
