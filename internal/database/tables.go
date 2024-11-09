package database

import "context"

func createTables(ctx context.Context) {
	createMovieTable := `
		CREATE TABLE IF NOT EXISTS movies (
			id				SERIAL PRIMARY KEY,
			title			TEXT NOT NULL,
			synopsis		TEXT NOT NULL,
			genre			TEXT[] NOT NULL,
			director		TEXT[] NOT NULL,
			"cast"			TEXT[] NOT NULL,
			release_date	TIMESTAMP,
			running_time	INTEGER NOT NULL,
			age_group		INTEGER NOT NULL,
			created_at		TIMESTAMP DEFAULT NOW(),
			updated_at		TIMESTAMP DEFAULT NOW(),
			deleted_at		TIMESTAMP NULL	
		)`

	_, err := Conn.Exec(ctx, createMovieTable)

	if err != nil {
		panic("Could not create movies table.")
	}

	createAddressTable := `
		CREATE TABLE IF NOT EXISTS address (
			id			SERIAL PRIMARY KEY,
			city		TEXT NOT NULL,
			street		TEXT NOT NULL,
			zip			TEXT NOT NULL,
			state		TEXT NOT NULL,
			country		TEXT NOT NULL,
			created_at	TIMESTAMP DEFAULT NOW(),
			updated_at	TIMESTAMP DEFAULT NOW(),
			deleted_at	TIMESTAMP NULL		
	)`

	_, err = Conn.Exec(ctx, createAddressTable)
	
	if err != nil {
		panic("Could not create address table.")
	}

	createTheaterTable := `
		CREATE TABLE IF NOT EXISTS theaters (
			id			SERIAL PRIMARY KEY,	
			name		TEXT NOT NULL UNIQUE,
			address_id	INTEGER REFERENCES address(id),	
			created_at	TIMESTAMP DEFAULT NOW(),
			updated_at	TIMESTAMP DEFAULT NOW(),
			deleted_at	TIMESTAMP NULL		
	)`

	_, err = Conn.Exec(ctx, createTheaterTable)

	if err != nil {
		panic("Could not create theater table.")
	}
}
