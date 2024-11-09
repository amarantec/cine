package theater

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/amarantec/cine/internal"
)

type TheaterRepository interface {
	ListTheaters(ctx context.Context) ([]internal.Theater, error)
	GetTheaterById(ctx context.Context, id uint) (internal.Theater, error)
	AddTheater(ctx context.Context, theater internal.Theater) (uint, error) 
}

type theaterRepository struct {
	Conn *pgxpool.Pool
}

func NewTheaterRepository(conn *pgxpool.Pool) TheaterRepository {
	return &theaterRepository{Conn: conn}
}

func (r *theaterRepository) ListTheaters(ctx context.Context) ([]internal.Theater, error) {
	rows, err := r.Conn.Query(
		ctx,
		`SELECT id, name, address_id, created_at, updated_at, deleted_at FROM theaters;`)

	if err != nil {
		return []internal.Theater{}, err
	}

	defer rows.Close()

	theaters := []internal.Theater{}
	for rows.Next() {
		t := internal.Theater{}
		if err := rows.Scan(
			&t.Id,
			&t.Name,
			&t.AddressId,
			&t.CreatedAt,
			&t.UpdatedAt,
			&t.DeletedAt); err != nil {
			return []internal.Theater{}, err
		}

		theaters = append(theaters, t)
	}

	if err := rows.Err(); err != nil {
		return []internal.Theater{}, err
	}

	return theaters, nil
}

func (r *theaterRepository) GetTheaterById(ctx context.Context, id uint) (internal.Theater, error) {
	theater := internal.Theater{Id: id}

	if err :=
		r.Conn.QueryRow(
			ctx,
			`SELECT name, address_id, created_at, updated_at, deleted_at
				 FROM theaters WHERE id = $1;`, id).Scan(&theater.Name, &theater.AddressId, &theater.CreatedAt, &theater.UpdatedAt, &theater.DeletedAt); err != nil {
		if err == pgx.ErrNoRows {
			return internal.Theater{}, nil
		}
		return internal.Theater{}, err
	}

	return theater, nil
}

func (r *theaterRepository) AddTheater(ctx context.Context, theater internal.Theater) (uint, error) {
	if err :=
		r.Conn.QueryRow(
			ctx,
			`INSERT INTO theaters (name, address_id) 
				VALUES ($1, $2) 
				RETURNING id;`, theater.Name, theater.AddressId).Scan(&theater.Id); err	!= nil {
					return internal.ZERO, err
				}
	
	return theater.Id, nil
}					
