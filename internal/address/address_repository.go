package address

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/amarantec/cine/internal"
	"time"
	"fmt"
)

type AddressRepository interface {
	InsertAddress(ctx context.Context, address internal.Address) (uint, error)
	GetAddress(ctx context.Context, id uint) (internal.Address, error)
	UpdateAddress(ctx context.Context, address internal.Address) (bool, error)
	DeleteAddress(ctx context.Context, id uint) (bool, error)
}

type addressRepository struct {
	Conn *pgxpool.Pool
}

func NewAddressRepository(conn *pgxpool.Pool) AddressRepository {
	return &addressRepository{Conn: conn}
}

func (r *addressRepository) InsertAddress(ctx context.Context, address internal.Address) (uint, error) {
	if err :=
		r.Conn.QueryRow(
			ctx,
			`INSERT INTO address (city, street, zip, state, country) 
				VALUES ($1, $2, $3, $4, $5) 
				RETURNING id;`, address.City, address.Street, address.ZIP,
								address.State, address.Country).Scan(&address.Id); err != nil {
									return internal.ZERO, err
				}

	return address.Id, nil
}

func (r *addressRepository) GetAddress(ctx context.Context, id uint) (internal.Address, error) {
	a := internal.Address{Id: id}

	if err :=
		r.Conn.QueryRow(
			ctx,
			`SELECT city, 
				street, 
				zip, 
				state, 
				country, 
				created_at, 
				updated_at 
				FROM address 
				WHERE id = $1 AND 
				deleted_at IS NULL`, id).Scan(&a.City, 
													&a.Street,
													&a.ZIP, 
													&a.State, 
													&a.Country, 
													&a.CreatedAt, 
													&a.UpdatedAt); err != nil {
													if err == pgx.ErrNoRows {
														return internal.Address{}, nil
													}
													return internal.Address{}, err
	}
	return a, nil
}

func (r *addressRepository) UpdateAddress(ctx context.Context, address internal.Address) (bool, error) {
	res, err :=
		r.Conn.Exec(	
			ctx,
			`UPDATE address SET city = $2, street = $3, zip = $4, state = $5, country = $6, updated_at = $7
				WHERE id = $1 AND deleted_at IS NULL;`, address.Id, address.City, address.Street,
														address.ZIP, address.State, address.Country, time.Now())
	
	if err != nil {
		return false, err
	}

	if res.RowsAffected() == 0 {
		fmt.Println("No rows affected")
		return false, nil
	} else {
		fmt.Printf("%d rows affected.\n", res.RowsAffected())
		return true, nil
	}
}	

func (r *addressRepository) DeleteAddress(ctx context.Context, id uint) (bool, error) {
	res, err :=
		r.Conn.Exec(
			ctx,
			`UPDATE	address SET deleted_at = $2 WHERE id = $1;`, id, time.Now())


	if err != nil {
		return false, err
	}

	if res.RowsAffected() == 0 {
		fmt.Println("No rows affected")
		return false, nil
	} else {
		fmt.Printf("%d rows affected.\n", res.RowsAffected())
		return true, nil
	}
}	
