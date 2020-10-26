package models

type SeatType int

const (
	Recliner SeatType = iota + 1
	Premium
	FrontRow
	Balcony
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
	SeatNumber int      `json:"seat_number"`
	Type       SeatType `json:"type"`

	// relations
	CinemaScreenID int          `json:"-"`
	CinemaScreen   CinemaScreen `json:"-"`
}
