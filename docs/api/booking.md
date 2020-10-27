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