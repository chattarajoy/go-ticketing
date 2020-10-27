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

access site on: http://localhost:4000/

# Documentation

* [API Docs](docs/api/index.md)
* [Database Design](docs/database_design.md)
* [Code Structure](docs/code_structure.md)

# TODO

- [ ] Add tests for more cases
- [ ] Add tests for helpers
- [ ] Fix MYSQL setup in github actions
- [ ] Add Authentication
- [ ] Add a user type field for theatre owners
- [ ] Add logger to each struct in services / api handlers
