# Ticketing (REST API for booking movie tickets)

## List of features

### Cinema Service
* Add Cinema (Cinema can be a hall or a multiplex)
* Add CinemaScreen (a screen in a hall or multiplex)

### Movies Service
* Add Movies (any movie that has been released)
* Add Movie Shows (add a show to screen from list of movies)
* List Movie Shows (list of shows across screens and halls)

###  Booking Service
* Show Movie Show seat status (show status of all bookings)
* Book Ticket with multiple seat selection (book multiple seats in a screen show)

# Running Locally using Docker

```bash
docker-compose up
```

access site on: https://localhost:4000/


# API Contracts and Responses

## Standard Response Format

```go
type Response struct {
	Success      bool        `json:"success"`
	StatusCode   int         `json:"status_code"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"error_message"`
}
```

## Cinema APIs

### `GET` `/cinemas` list all cinemas and their properties (city, seats etc)

Output

```go
type ListCinemasOutput struct {
	Cinemas []models.Cinema `json:"cinemas"`
}
```

Sample Response 

```json
{
    "success": true,
    "status_code": 200,
    "data": {
        "cinemas": [
            {
                "id": 1,
                "name": "Cinema from Postman 4",
                "screens": [
                    {
                        "id": 1,
                        "name": "Screen 1",
                        "seats": [
                            {
                                "id": 1,
                                "seat_number": 1,
                                "type": "RECLINER"
                            },
                            {
                                "id": 2,
                                "seat_number": 2,
                                "type": "RECLINER"
                            }
                        ]
                    },
                    {
                        "id": 2,
                        "name": "Screen 1",
                        "seats": [
                            {
                                "id": 4,
                                "seat_number": 1,
                                "type": "RECLINER"
                            },
                            {
                                "id": 5,
                                "seat_number": 2,
                                "type": "RECLINER"
                            }
                        ]
                    },
                    {
                        "id": 3,
                        "name": "Screen 1",
                        "seats": [
                            {
                                "id": 6,
                                "seat_number": 1,
                                "type": "RECLINER"
                            },
                            {
                                "id": 7,
                                "seat_number": 2,
                                "type": "RECLINER"
                            },
                            {
                                "id": 8,
                                "seat_number": 3,
                                "type": "RECLINER"
                            },
                            {
                                "id": 9,
                                "seat_number": 4,
                                "type": "RECLINER"
                            }
                        ]
                    },
                    {
                        "id": 4,
                        "name": "Screen 1",
                        "seats": [
                            {
                                "id": 10,
                                "seat_number": 1,
                                "type": "RECLINER"
                            },
                            {
                                "id": 11,
                                "seat_number": 2,
                                "type": "RECLINER"
                            },
                            {
                                "id": 12,
                                "seat_number": 3,
                                "type": "RECLINER"
                            },
                            {
                                "id": 13,
                                "seat_number": 4,
                                "type": "RECLINER"
                            }
                        ]
                    }
                ],
                "city": {
                    "id": 1,
                    "name": "Bengaluru",
                    "zip_code": "560068"
                }
            }
        ]
    },
    "error_message": ""
}
```

### `POST` `/cinema` Add a new cinema

Input 

```go
type AddCinemaInput struct {
	CinemaName string `json:"cinema_name"`
	CityID     int    `json:"city_id"`
}
```

Output 

```go
type AddCinemaOutput struct {
	Cinema models.Cinema `json:"cinema"`
}
```

Sample Response 

```json
{
    "success": true,
    "status_code": 200,
    "data": {
        "cinema": {
            "id": 2,
            "name": "Cinema from Postman 4",
            "screens": null,
            "city": {
                "id": 0,
                "name": "",
                "zip_code": ""
            }
        }
    },
    "error_message": ""
}
```


### `POST` `/screen`  Add a screen to Cinema

Input 

```go
type AddCinemaScreenInput struct {
	CinemaID   int         `json:"cinema_id"`
	ScreenName string      `json:"screen_name"`
	Seats      []*SeatInfo `json:"seats"`
}
```

Output 

```go
type AddCinemaScreenOutput struct {
	CinemaScreen models.CinemaScreen `json:"cinema_screen"`
}
```

Sample Response 

```json
{
    "success": true,
    "status_code": 200,
    "data": {
        "cinema_screen": {
            "id": 5,
            "name": "Screen 1",
            "seats": [
                {
                    "id": 14,
                    "seat_number": 1,
                    "type": "RECLINER"
                },
                {
                    "id": 15,
                    "seat_number": 2,
                    "type": "RECLINER"
                },
                {
                    "id": 16,
                    "seat_number": 3,
                    "type": "RECLINER"
                },
                {
                    "id": 17,
                    "seat_number": 4,
                    "type": "RECLINER"
                }
            ]
        }
    },
    "error_message": ""
}
```

## Movie APIs

### `POST` `/movie` Add a new Movie

Input

```go
type AddMovieInput struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Duration    time.Duration `json:"duration"`
}
```

Output

```go
type AddMovieOutput struct {
	Movie models.Movie `json:"movie"`
}
```

Sample Response

```json
{
    "success": true,
    "status_code": 200,
    "data": {
        "movie": {
            "id": 2,
            "name": "New Movie",
            "description": "It's a new movie",
            "duration": 60000,
            "shows": null
        }
    },
    "error_message": ""
}
```


### `POST` `/show` Add a new show of movie to cinema screen

Input

```go
type AddMovieShowInput struct {
	MovieID        int       `json:"movie_id"`
	CinemaScreenID int       `json:"cinema_screen_id"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
}
```

Output

```go
type AddMovieShowOutput struct {
	Show models.MovieShow
}
```

Sample Response

```json
{
    "success": true,
    "status_code": 200,
    "data": {
        "Show": {
            "id": 3,
            "start_time": "2020-11-03T01:34:05+05:30",
            "end_time": "2020-11-03T02:34:05+05:30",
            "CinemaScreen": {
                "id": 0,
                "name": "",
                "seats": null
            },
            "Movie": {
                "id": 1,
                "name": "New Movie",
                "description": "It's a new movie",
                "duration": 60000,
                "shows": null
            },
            "bookings": [],
            "seats": [
                {
                    "status": "AVAILABLE",
                    "cinema_seat": {
                        "id": 0,
                        "seat_number": 0,
                        "type": ""
                    }
                },
                {
                    "status": "AVAILABLE",
                    "cinema_seat": {
                        "id": 0,
                        "seat_number": 0,
                        "type": ""
                    }
                }
            ]
        }
    },
    "error_message": ""
}
```

### `GET` `/show` Get a show and all its movie, seats etc

Input

```go
type GetMovieShowInput struct {
	ShowID int `json:"show_id"`
} 
```


Output

```go
type GetMovieShowOutput struct {
	Show models.MovieShow `json:"show"`
}
```

Sample Response 

```json
{
    "success": true,
    "status_code": 200,
    "data": {
        "show": {
            "id": 2,
            "start_time": "2020-11-02T23:34:05+05:30",
            "end_time": "2020-11-03T00:34:05+05:30",
            "CinemaScreen": {
                "id": 4,
                "name": "Screen 1",
                "seats": null
            },
            "Movie": {
                "id": 1,
                "name": "New Movie",
                "description": "It's a new movie",
                "duration": 60000,
                "shows": null
            },
            "bookings": [
                {
                    "id": 16,
                    "seat_count": 2,
                    "status": "CONFIRMED",
                    "user": {
                        "id": 0,
                        "Name": "",
                        "Email": "",
                        "Bookings": null
                    },
                    "movie_show": {
                        "id": 0,
                        "start_time": "0001-01-01T00:00:00Z",
                        "end_time": "0001-01-01T00:00:00Z",
                        "CinemaScreen": {
                            "id": 0,
                            "name": "",
                            "seats": null
                        },
                        "Movie": {
                            "id": 0,
                            "name": "",
                            "description": "",
                            "duration": 0,
                            "shows": null
                        },
                        "bookings": null,
                        "seats": null
                    },
                    "seats": null
                },
                {
                    "id": 17,
                    "seat_count": 2,
                    "status": "FAILED",
                    "user": {
                        "id": 0,
                        "Name": "",
                        "Email": "",
                        "Bookings": null
                    },
                    "movie_show": {
                        "id": 0,
                        "start_time": "0001-01-01T00:00:00Z",
                        "end_time": "0001-01-01T00:00:00Z",
                        "CinemaScreen": {
                            "id": 0,
                            "name": "",
                            "seats": null
                        },
                        "Movie": {
                            "id": 0,
                            "name": "",
                            "description": "",
                            "duration": 0,
                            "shows": null
                        },
                        "bookings": null,
                        "seats": null
                    },
                    "seats": null
                },
                {
                    "id": 18,
                    "seat_count": 1,
                    "status": "CONFIRMED",
                    "user": {
                        "id": 0,
                        "Name": "",
                        "Email": "",
                        "Bookings": null
                    },
                    "movie_show": {
                        "id": 0,
                        "start_time": "0001-01-01T00:00:00Z",
                        "end_time": "0001-01-01T00:00:00Z",
                        "CinemaScreen": {
                            "id": 0,
                            "name": "",
                            "seats": null
                        },
                        "Movie": {
                            "id": 0,
                            "name": "",
                            "description": "",
                            "duration": 0,
                            "shows": null
                        },
                        "bookings": null,
                        "seats": null
                    },
                    "seats": null
                }
            ],
            "seats": [
                {
                    "status": "BOOKED",
                    "cinema_seat": {
                        "id": 0,
                        "seat_number": 0,
                        "type": ""
                    }
                },
                {
                    "status": "AVAILABLE",
                    "cinema_seat": {
                        "id": 0,
                        "seat_number": 0,
                        "type": ""
                    }
                },
                {
                    "status": "BOOKED",
                    "cinema_seat": {
                        "id": 0,
                        "seat_number": 0,
                        "type": ""
                    }
                },
                {
                    "status": "BOOKED",
                    "cinema_seat": {
                        "id": 0,
                        "seat_number": 0,
                        "type": ""
                    }
                }
            ]
        }
    },
    "error_message": ""
}
```

## Booking APIs

### `GET` `/bookings` List of all the bookings

Output

```go
type ListBookingsOutput struct {
	Bookings []models.Booking `json:"bookings"`
}
```

Sample Response

```json
{
    "success": true,
    "status_code": 200,
    "data": {
        "bookings": [
            {
                "id": 1,
                "seat_count": 2,
                "status": "PENDING",
                "user": {
                    "id": 1,
                    "Name": "Joy",
                    "Email": "joylal4896@gmail.com",
                    "Bookings": null
                },
                "movie_show": {
                    "id": 1,
                    "start_time": "2020-11-02T23:34:05+05:30",
                    "end_time": "2020-11-03T00:34:05+05:30",
                    "CinemaScreen": {
                        "id": 0,
                        "name": "",
                        "seats": null
                    },
                    "Movie": {
                        "id": 0,
                        "name": "",
                        "description": "",
                        "duration": 0,
                        "shows": null
                    },
                    "bookings": null,
                    "seats": null
                },
                "seats": []
            },
            {
                "id": 14,
                "seat_count": 2,
                "status": "FAILED",
                "user": {
                    "id": 1,
                    "Name": "Joy",
                    "Email": "joylal4896@gmail.com",
                    "Bookings": null
                },
                "movie_show": {
                    "id": 1,
                    "start_time": "2020-11-02T23:34:05+05:30",
                    "end_time": "2020-11-03T00:34:05+05:30",
                    "CinemaScreen": {
                        "id": 0,
                        "name": "",
                        "seats": null
                    },
                    "Movie": {
                        "id": 0,
                        "name": "",
                        "description": "",
                        "duration": 0,
                        "shows": null
                    },
                    "bookings": null,
                    "seats": null
                },
                "seats": []
            },
            {
                "id": 15,
                "seat_count": 2,
                "status": "FAILED",
                "user": {
                    "id": 1,
                    "Name": "Joy",
                    "Email": "joylal4896@gmail.com",
                    "Bookings": null
                },
                "movie_show": {
                    "id": 1,
                    "start_time": "2020-11-02T23:34:05+05:30",
                    "end_time": "2020-11-03T00:34:05+05:30",
                    "CinemaScreen": {
                        "id": 0,
                        "name": "",
                        "seats": null
                    },
                    "Movie": {
                        "id": 0,
                        "name": "",
                        "description": "",
                        "duration": 0,
                        "shows": null
                    },
                    "bookings": null,
                    "seats": null
                },
                "seats": []
            },
            {
                "id": 16,
                "seat_count": 2,
                "status": "CONFIRMED",
                "user": {
                    "id": 1,
                    "Name": "Joy",
                    "Email": "joylal4896@gmail.com",
                    "Bookings": null
                },
                "movie_show": {
                    "id": 2,
                    "start_time": "2020-11-02T23:34:05+05:30",
                    "end_time": "2020-11-03T00:34:05+05:30",
                    "CinemaScreen": {
                        "id": 0,
                        "name": "",
                        "seats": null
                    },
                    "Movie": {
                        "id": 0,
                        "name": "",
                        "description": "",
                        "duration": 0,
                        "shows": null
                    },
                    "bookings": null,
                    "seats": null
                },
                "seats": [
                    {
                        "status": "BOOKED",
                        "cinema_seat": {
                            "id": 0,
                            "seat_number": 0,
                            "type": ""
                        }
                    },
                    {
                        "status": "BOOKED",
                        "cinema_seat": {
                            "id": 0,
                            "seat_number": 0,
                            "type": ""
                        }
                    }
                ]
            },
            {
                "id": 17,
                "seat_count": 2,
                "status": "FAILED",
                "user": {
                    "id": 1,
                    "Name": "Joy",
                    "Email": "joylal4896@gmail.com",
                    "Bookings": null
                },
                "movie_show": {
                    "id": 2,
                    "start_time": "2020-11-02T23:34:05+05:30",
                    "end_time": "2020-11-03T00:34:05+05:30",
                    "CinemaScreen": {
                        "id": 0,
                        "name": "",
                        "seats": null
                    },
                    "Movie": {
                        "id": 0,
                        "name": "",
                        "description": "",
                        "duration": 0,
                        "shows": null
                    },
                    "bookings": null,
                    "seats": null
                },
                "seats": []
            },
            {
                "id": 18,
                "seat_count": 1,
                "status": "CONFIRMED",
                "user": {
                    "id": 1,
                    "Name": "Joy",
                    "Email": "joylal4896@gmail.com",
                    "Bookings": null
                },
                "movie_show": {
                    "id": 2,
                    "start_time": "2020-11-02T23:34:05+05:30",
                    "end_time": "2020-11-03T00:34:05+05:30",
                    "CinemaScreen": {
                        "id": 0,
                        "name": "",
                        "seats": null
                    },
                    "Movie": {
                        "id": 0,
                        "name": "",
                        "description": "",
                        "duration": 0,
                        "shows": null
                    },
                    "bookings": null,
                    "seats": null
                },
                "seats": [
                    {
                        "status": "BOOKED",
                        "cinema_seat": {
                            "id": 0,
                            "seat_number": 0,
                            "type": ""
                        }
                    }
                ]
            }
        ]
    },
    "error_message": ""
}
```

### `POST` `/book` Book a list of seats in a show

Input

```go
type BookSeatsInput struct {
	ShowID      int             `json:"show_id"`
	SeatNumbers []int           `json:"seat_numbers"`
	UserID      int             `json:"user_id"`
	SeatType    models.SeatType `json:"seat_type"`
}
```


Output

```go
type BookSeatsOutput struct {
	Booking models.Booking `json:"booking"`
}
```

Sample Response 

```json
{
    "success": true,
    "status_code": 200,
    "data": {
        "booking": {
            "id": 21,
            "seat_count": 1,
            "status": "CONFIRMED",
            "user": {
                "id": 1,
                "Name": "Joy",
                "Email": "joylal4896@gmail.com",
                "Bookings": null
            },
            "movie_show": {
                "id": 3,
                "start_time": "2020-11-03T01:34:05+05:30",
                "end_time": "2020-11-03T02:34:05+05:30",
                "CinemaScreen": {
                    "id": 0,
                    "name": "",
                    "seats": null
                },
                "Movie": {
                    "id": 0,
                    "name": "",
                    "description": "",
                    "duration": 0,
                    "shows": null
                },
                "bookings": null,
                "seats": null
            },
            "seats": [
                {
                    "status": "BOOKED",
                    "cinema_seat": {
                        "id": 0,
                        "seat_number": 0,
                        "type": ""
                    }
                }
            ]
        }
    },
    "error_message": ""
}
```