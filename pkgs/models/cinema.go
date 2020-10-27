package models

type SeatType string

const (
	Recliner SeatType = "RECLINER"
	Premium  SeatType = "PREMIUM"
	FrontRow SeatType = "FRONT"
	Balcony  SeatType = "BALCONY"
)

// TODO: add unique index on ZIP Code
type City struct {
	Model
	Name    string `json:"name"`
	ZipCode string `json:"zip_code"`
}

// Cinema: a cinema hall or a multiplex
type Cinema struct {
	Model
	Name string `json:"name"`

	// relations
	CinemaScreens []CinemaScreen `json:"screens" gorm:"foreignkey:CinemaID"`
	CityID        int            `json:"-"`
	City          City           `json:"city"`
}

// CinemaScreen: cinema screen which denotes a specific screen on the of Cinema
type CinemaScreen struct {
	Model
	Name string `json:"name"`

	// relations
	CinemaID    int          `json:"-"`
	Cinema      Cinema       `json:"-"`
	CinemaSeats []CinemaSeat `json:"seats" gorm:"foreignkey:CinemaScreenID"`
}

// CinemaSeat: all the seats in a cinema screen
type CinemaSeat struct {
	Model
	SeatNumber int      `json:"seat_number" gorm:"index:unique_seat_per_cinema_screen_and_type,unique"`
	Type       SeatType `json:"type" gorm:"index:unique_seat_per_cinema_screen_and_type,unique"`

	// relations
	CinemaScreenID int          `json:"-" gorm:"index:unique_seat_per_cinema_screen_and_type,unique"`
	CinemaScreen   CinemaScreen `json:"-"`
}
