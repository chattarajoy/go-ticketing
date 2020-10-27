
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