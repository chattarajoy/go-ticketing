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