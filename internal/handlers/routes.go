package handlers

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/amarantec/cine/internal/movie"
	"gitlab.com/amarantec/cine/internal/theater"
	"gitlab.com/amarantec/cine/internal/address"
	"gitlab.com/amarantec/cine/internal/room"
     httpSwagger "github.com/swaggo/http-swagger"
     _ "gitlab.com/amarantec/cine/docs/swagger"
)

func SetRoutes(conn *pgxpool.Pool) *http.ServeMux {
	m := http.NewServeMux()

	// Movie dependency injection
	movieRepository := movie.NewMovieRepository(conn)
	movieService := movie.MovieService(movieRepository)
	movieHandler := NewMovieHandler(movieService)

	// Theater dependency injection
	theaterRepository := theater.NewTheaterRepository(conn)
	theaterService := theater.NewTheaterService(theaterRepository)
	theaterHandler := NewTheaterHandler(theaterService)

	// Address dependency injection
	addressRepository := address.NewAddressRepository(conn)
	addressService := address.NewAddressService(addressRepository)	
	addressHandler := NewAddressHandler(addressService)

	// Cine Room dependency injection
	roomRepository := room.NewRoomRepository(conn)
	roomService := room.NewRoomService(roomRepository)
	roomHandler := NewRoomHandler(roomService)
    
	/*
		ROUTES
	*/

	m.HandleFunc("/list-movies", movieHandler.listMovies)
	m.HandleFunc("/get-movie-by-id/{id}", movieHandler.getMovieById)
	m.HandleFunc("/add-movie", movieHandler.addMovie)
	m.HandleFunc("/get-movies-by-genre/{genre}", movieHandler.getMoviesByGenre)
    m.HandleFunc("/update-movie", movieHandler.updateMovie)
    m.HandleFunc("/delete-movie/{movieId}", movieHandler.deleteMovie)

	m.HandleFunc("/list-theaters", theaterHandler.listTheaters)
	m.HandleFunc("/get-theater-by-id/{id}", theaterHandler.getTheaterById)
	m.HandleFunc("/add-theater", theaterHandler.addTheater)

	m.HandleFunc("/insert-address", addressHandler.insertAddress)
	m.HandleFunc("/get-address/{id}", addressHandler.getAddress)
	m.HandleFunc("/update-address", addressHandler.updateAddress)
	m.HandleFunc("/delete-address/{id}", addressHandler.deleteAddress)
	
	m.HandleFunc("/list-room", roomHandler.listRooms)
	m.HandleFunc("/get-room-by-id/{theaterId}", roomHandler.getRoomById)
	m.HandleFunc("/list-available-room-seats/{theaterId}/{roomNumber}", roomHandler.listAvailableRoomSeats)
 


    m.HandleFunc("/swagger/", httpSwagger.WrapHandler) // Access the Swagger UI
	return m
}
