package models

type User struct {
	Model
	Name  string
	Email string

	// relations
	Bookings []Booking `gorm:"foreignKey:UserID"`
}
