package models

type BookingStatus int
type SeatStatus int

const (
	BookingConfirmed BookingStatus = iota + 1
	BookingCancelled
	BookingFailed
	BookingPending

	SeatAvailable SeatStatus = iota + 1
	SeatBooked
)

type Booking struct {
	Model
	SeatCount int           `json:"seat_count"`
	Status    BookingStatus `json:"status"`

	// relations
	UserID      int       `json:"-"`
	User        User      `json:"user"`
	MovieShowID int       `json:"-"`
	MovieShow   MovieShow `json:"movie_show"`
}

type BookingSeat struct {
	Status       SeatStatus `json:"status"`
	MovieShowID  int        `json:"-" gorm:"index:unique_seat_per_show,unique"`
	CinemaSeatID int        `json:"-" gorm:"index:unique_seat_per_show,unique"`
	BookingID    int        `json:"-"`

	// relations
	MovieShow  MovieShow  `json:"movie_show"`
	CinemaSeat CinemaSeat `json:"cinema_seat"`
	Booking    Booking    `json:"booking"`
}
