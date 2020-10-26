package models

type SeatType int

const (
	Recliner SeatType = iota + 1
	Premium
	FrontRow
	Balcony
)

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
	CinemaScreens []CinemaScreen `json:"cinema_screens" gorm:"foreignkey:CinemaID"`
	CityID        int            `json:"-"`
	City          City           `json:"city"`
}

// CinemaScreen: cinema screen which denotes a specific screen on the of Cinema
type CinemaScreen struct {
	Model
	Name string `json:"name"`

	// relations
	CinemaID    int          `json:"-"`
	Cinema      Cinema       `json:"cinema"`
	CinemaSeats []CinemaSeat `gorm:"foreignKey:CinemaScreenID" json:"cinema_seats"`
}

// CinemaSeat: all the seats in a cinema screen
type CinemaSeat struct {
	Model
	SeatNumber int      `json:"seat_number"`
	Type       SeatType `json:"type"`

	// relations
	CinemaScreenID int          `json:"-"`
	CinemaScreen   CinemaScreen `json:"cinema_screen"`
}
