package booking

import (
	"reflect"
	"testing"

	"gorm.io/gorm"

	"commerceiq.ai/ticketing/internal/cache"
	"commerceiq.ai/ticketing/internal/testhelpers"
	"commerceiq.ai/ticketing/pkgs/models"
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

func TestService_BookSeats(t *testing.T) {
	db := testhelpers.SetupDB()
	cacheOb := cache.NewCache(cache.InMemoryCache)

	type fields struct {
		db    *gorm.DB
		cache cache.Cache
	}
	type args struct {
		inp *BookSeatsInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *BookSeatsOutput
		wantErr bool
	}{
		{
			name: "TestBookSeats",
			fields: fields{
				db:    db,
				cache: cacheOb,
			},
			args: args{inp: &BookSeatsInput{
				ShowID:      1,
				SeatNumbers: []int{1, 2},
				UserID:      1,
				SeatType:    models.Recliner,
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				db:    tt.fields.db,
				cache: tt.fields.cache,
			}
			got, err := s.BookSeats(tt.args.inp)
			if (err != nil) != tt.wantErr {
				t.Errorf("BookSeats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BookSeats() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ListBookings(t *testing.T) {
	db := testhelpers.SetupDB()
	cacheOb := cache.NewCache(cache.InMemoryCache)

	type fields struct {
		db    *gorm.DB
		cache cache.Cache
	}
	tests := []struct {
		name    string
		fields  fields
		want    *ListBookingsOutput
		wantErr bool
	}{
		{
			name: "TestListBookings",
			fields: fields{
				db:    db,
				cache: cacheOb,
			},
			want: &ListBookingsOutput{
				Bookings: []models.Booking{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				db:    tt.fields.db,
				cache: tt.fields.cache,
			}
			got, err := s.ListBookings()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListBookings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListBookings() got = %v, want %v", got, tt.want)
			}
		})
	}
}
