package models

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
)

type BookingStatus string
type SeatStatus string

const (
	BookingConfirmed BookingStatus = "CONFIRMED"
	BookingCancelled BookingStatus = "CANCELLED"
	BookingFailed    BookingStatus = "FAILED"
	BookingPending   BookingStatus = "PENDING"

	SeatAvailable SeatStatus = "AVAILABLE"
	SeatBooked    SeatStatus = "BOOKED"
)

type Booking struct {
	Model
	SeatCount int           `json:"seat_count"`
	Status    BookingStatus `json:"status"`

	// relations
	UserID      int           `json:"-"`
	User        User          `json:"user"`
	MovieShowID int           `json:"-"`
	MovieShow   MovieShow     `json:"movie_show"`
	Seats       []BookingSeat `json:"seats" gorm:"foreignKey:BookingID"`
}

func (b *Booking) BeforeSave(db *gorm.DB) (err error) {
	b.MovieShow = MovieShow{}
	return
}

// TODO: Break into smaller logical methods
func (b *Booking) BookSeats(db *gorm.DB, seatNumbers []int, seatType SeatType) error {
	var seats []int
	result := db.Model(&CinemaSeat{}).
		Where("cinema_screen_id = ? AND seat_number IN ? AND type = ? ", b.MovieShow.CinemaScreenID, seatNumbers, seatType).
		Pluck("ID", &seats)

	if len(seatNumbers) != len(seats) {
		return fmt.Errorf("%d seat numbers are invalid in the list: %v", len(seatNumbers)-len(seats), seatNumbers)
	}

	result = db.Model(BookingSeat{}).
		Where("cinema_seat_id IN ?", seats).
		Where("status = ? ", SeatAvailable).
		Where("movie_show_id = ? ", b.MovieShow.ID).
		Updates(map[string]interface{}{"status": SeatBooked, "booking_id": b.ID})

	if result.RowsAffected != int64(len(seats)) {
		errorMsg := "race condition while booking, some seats got booked already"
		revert := db.Model(BookingSeat{}).
			Where("cinema_seat_id IN ?", seats).
			Where("status = ?", SeatBooked).
			Where("booking_id = ?", b.ID).
			Updates(map[string]interface{}{"status": SeatAvailable, "booking_id": 0})
		if revert.Error != nil {
			errorMsg += "Error, while reverting, " + revert.Error.Error()
		}
		return fmt.Errorf(errorMsg)
	}
	return nil
}

func (b *Booking) Fail() {
	b.Status = BookingFailed
}

func (b *Booking) Confirm() {
	b.Status = BookingConfirmed
}

type BookingSeat struct {
	Status       SeatStatus    `json:"status"`
	MovieShowID  int           `json:"-" gorm:"index:unique_seat_per_show,unique"`
	CinemaSeatID int           `json:"-" gorm:"index:unique_seat_per_show,unique"`
	BookingID    sql.NullInt64 `json:"-"`

	// relations
	MovieShow  MovieShow  `json:"-"`
	CinemaSeat CinemaSeat `json:"cinema_seat"`
	Booking    Booking    `json:"-"`
}
