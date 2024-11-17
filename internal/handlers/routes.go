package handlers

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/amarantec/cine/internal/movie"
	"gitlab.com/amarantec/cine/internal/theater"
	"gitlab.com/amarantec/cine/internal/address"
	"gitlab.com/amarantec/cine/internal/room"
)

func SetRoutes(conn *pgxpool.Pool) *http.ServeMux {
	m := http.NewServeMux()

	// Movie dependency injection
    movieMux := http.NewServeMux()
	movieRepository := movie.NewMovieRepository(conn)
	movieService := movie.MovieService(movieRepository)
	movieHandler := NewMovieHandler(movieService)

	// Theater dependency injection
    theaterMux := http.NewServeMux()
	theaterRepository := theater.NewTheaterRepository(conn)
	theaterService := theater.NewTheaterService(theaterRepository)
	theaterHandler := NewTheaterHandler(theaterService)

	// Address dependency injection
	addressMux := http.NewServeMux()
	addressRepository := address.NewAddressRepository(conn)
	addressService := address.NewAddressService(addressRepository)	
	addressHandler := NewAddressHandler(addressService)

	// Cine Room dependency injection
	roomMux := http.NewServeMux()
	roomRepository := room.NewRoomRepository(conn)
	roomService := room.NewRoomService(roomRepository)
	roomHandler := NewRoomHandler(roomService)
	/*
		ROUTES
	*/

	movieMux.HandleFunc("/list-movies", movieHandler.listMovies)
	movieMux.HandleFunc("/get-movie-by-id/{id}", movieHandler.getMovieById)
	movieMux.HandleFunc("/add-movie", movieHandler.addMovie)
	movieMux.HandleFunc("/get-movies-by-genre/{genre}", movieHandler.getMoviesByGenre)


	theaterMux.HandleFunc("/list-theaters", theaterHandler.listTheaters)
	theaterMux.HandleFunc("/get-theater-by-id/{id}", theaterHandler.getTheaterById)
	theaterMux.HandleFunc("/add-theater", theaterHandler.addTheater)

	addressMux.HandleFunc("/insert-address", addressHandler.insertAddress)
	addressMux.HandleFunc("/get-address/{id}", addressHandler.getAddress)
	addressMux.HandleFunc("/update-address", addressHandler.updateAddress)
	addressMux.HandleFunc("/delete-address/{id}", addressHandler.deleteAddress)
	
	roomMux.HandleFunc("/list-room", roomHandler.listRooms)
	roomMux.HandleFunc("/get-room-by-id/{theaterId}", roomHandler.getRoomById)
	roomMux.HandleFunc("/list-available-room-seats/{theaterId}/{roomNumber}", roomHandler.listAvailableRoomSeats)
 
    m.Handle("/movies/", http.StripPrefix("/movies", movieMux))
    m.Handle("/theaters/", http.StripPrefix("/theaters", theaterMux))
	m.Handle("/address/", http.StripPrefix("/address", addressMux))
	m.Handle("/room/", http.StripPrefix("/room", roomMux))

	return m
}
