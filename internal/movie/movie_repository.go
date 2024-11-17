package movie

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/amarantec/cine/internal"
)

type MovieRepository interface {
	ListMovies(ctx context.Context) ([]internal.Movie, error)
	GetMovieById(ctx context.Context, id uint) (internal.Movie, error)
	AddMovie(ctx context.Context, movie internal.Movie) (uint, error)
	GetMoviesByGenre(ctx context.Context, genre string) ([]internal.Movie, error)
}

type movieRepository struct {
	Conn *pgxpool.Pool
}

func NewMovieRepository(conn *pgxpool.Pool) MovieRepository {
	return &movieRepository{Conn: conn}
}

func (r *movieRepository) ListMovies(ctx context.Context) ([]internal.Movie, error) {

	rows, err := r.Conn.Query(
		ctx,
		`SELECT id,
			title, 
			synopsis, 
			genre,
			director, 
			"cast", 
			release_date, 
			running_time, 
			age_group,
			created_at,
			updated_at FROM movies WHERE deleted_at IS NULL;`)

	if err != nil {
		return []internal.Movie{}, err
	}
	defer rows.Close()

	movies := []internal.Movie{}
	for rows.Next() {
		m := internal.Movie{}
		if err := rows.Scan(
			&m.Id,
			&m.Title,
			&m.Synopsis,
			&m.Genre,
			&m.Director,
			&m.Cast,
			&m.ReleaseDate,
			&m.RunningTime,
			&m.AgeGroup,
			&m.CreatedAt,
			&m.UpdatedAt); err != nil {
			return []internal.Movie{}, err
		}
		movies = append(movies, m)
	}

	if err := rows.Err(); err != nil {
		return []internal.Movie{}, err
	}

	return movies, nil
}

func (r *movieRepository) GetMovieById(ctx context.Context, id uint) (internal.Movie, error) {
	movie := internal.Movie{Id: id}

	if err :=
		r.Conn.QueryRow(
			ctx,
			`SELECT title, 
				synopsis, 
				genre,
				director, 
				"cast", 
				release_date, 
				running_time, 
				age_group,
				created_at,
				updated_at
				FROM movies WHERE id = $1 AND deleted_at IS NULL;`, id).Scan(&movie.Title,
			&movie.Synopsis, &movie.Genre, &movie.Director, &movie.Cast, 
			&movie.ReleaseDate, &movie.RunningTime, &movie.AgeGroup,
			&movie.CreatedAt, &movie.UpdatedAt); err != nil {
		if err == pgx.ErrNoRows {
			return internal.Movie{}, nil
		}
		return internal.Movie{}, err
	}
	return movie, nil
}

func (r *movieRepository) AddMovie(ctx context.Context, movie internal.Movie) (uint, error) {
	if err :=
		r.Conn.QueryRow(
			ctx,
			`INSERT INTO movies (title, synopsis, genre, director, "cast", release_date, running_time, age_group) VALUES
			($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`, 
				movie.Title, 
				movie.Synopsis, 
				movie.Genre, 
				movie.Director, 
				movie.Cast, 
				movie.ReleaseDate, 
				movie.RunningTime, 
				movie.AgeGroup).Scan(&movie.Id); err != nil {
					return internal.ZERO, err
	}

	return movie.Id, nil
}

func (r *movieRepository) GetMoviesByGenre(ctx context.Context, genre string) ([]internal.Movie, error) {
	rows, err :=
		r.Conn.Query(
			ctx,
			`SELECT id,
				title,
				synopsis,
				genre,
				director,
				"cast",
				release_date,
				running_time,
				age_group FROM movies
				WHERE $1 = ANY(genre) AND deleted_at IS NULL;`, genre)

	if err != nil {
		return []internal.Movie{}, err
	}

	defer rows.Close()

	movies := []internal.Movie{}

	for rows.Next() {
		m := internal.Movie{}
		if err := rows.Scan(
			&m.Id,
			&m.Title,
			&m.Synopsis,
			&m.Genre,
			&m.Director,
			&m.Cast,
			&m.ReleaseDate,
			&m.RunningTime,
			&m.AgeGroup); err != nil {
			return []internal.Movie{}, err
		}
		movies = append(movies, m)
	}

	if err := rows.Err(); err != nil {
		return []internal.Movie{}, err
	}

	return movies, nil
}
