package testhelpers

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"commerceiq.ai/ticketing/pkgs/models"
)

func SetupDB() *gorm.DB { // or *gorm.DB
	dsn := "user:user@tcp(127.0.0.1:3306)/ticketing_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error connecting to DB: ", err.Error())
		os.Exit(1)
	}
	_ = db.AutoMigrate(&models.Booking{},
		&models.BookingSeat{},
		&models.Movie{},
		&models.MovieShow{},
		&models.User{},
		&models.City{},
		&models.Cinema{},
		&models.CinemaScreen{},
		&models.CinemaSeat{},
	)
	return db
}
