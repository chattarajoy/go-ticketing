package main

import (
	"github.com/chattarajoy/go-ticketing/cmd"
)

func main() {
	// dsn := "root:@tcp(127.0.0.1:3306)/commerceiq?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println("error connecting to DB: ", err.Error())
	// 	os.Exit(1)
	// }
	// _ = db.AutoMigrate(&models.Booking{},
	// 	&models.BookingSeat{},
	// 	&models.Movie{},
	// 	&models.MovieShow{},
	// 	&models.User{},
	// 	&models.City{},
	// 	&models.Cinema{},
	// 	&models.CinemaScreen{},
	// 	&models.CinemaSeat{},
	// )
	//
	// svc := cinema.NewService(db, nil)
	// resp := svc.AddCinema("test new cinema", 1)
	// if !resp.Success {
	// 	fmt.Println("error adding cinema: ", resp.ErrorMessage)
	// 	os.Exit(1)
	// }
	// fmt.Println("add", resp.StatusCode, string(resp.Data))
	//
	// resp = svc.ListCinemas()
	// if !resp.Success {
	// 	fmt.Println("error listing cinemas: ", resp.ErrorMessage)
	// 	os.Exit(1)
	// }
	// fmt.Println("list", resp.StatusCode, string(resp.Data))

	cmd.Execute()
}
