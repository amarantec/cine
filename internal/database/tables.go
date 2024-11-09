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
			age_group		INTEGER NOT NULL

		)`

	_, err := Conn.Exec(ctx, createMovieTable)

	if err != nil {
		panic("Could not create movies table.")
	}
}
