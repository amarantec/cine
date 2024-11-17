package room

import (
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/jackc/pgx/v5"
    "gitlab.com/amarantec/cine/internal"
    "context"
)

type RoomRepository interface {
	ListRooms(ctx context.Context, theatherId uint) ([]internal.CineRoom, error)
	GetRoomById(ctx context.Context, theaterId, roomId uint) (internal.CineRoom, error)
	ListAvailableRoomSeats(ctx context.Context, theaterId uint, room string) ([]internal.Seat, error)
}

type roomRepository struct {
    Conn *pgxpool.Pool
}

func NewRoomRepository(conn *pgxpool.Pool) RoomRepository {
    return &roomRepository{Conn: conn}
}


func (r *roomRepository) ListRooms(ctx context.Context, theaterId uint) ([]internal.CineRoom, error) {
    rows, err :=
        r.Conn.Query(
            ctx,
            `SELECT id, number, created_at, updated_at
                FROM rooms WHERE 
                theater_id = $1
                deleted_at IS NULL;`, theaterId)

    if err != nil {
        return []internal.CineRoom{}, err
    }

    defer rows.Close()

    rooms := []internal.CineRoom{}
    for rows.Next() {
        cr := internal.CineRoom{}
        if err := rows.Scan(
            &cr.Id,
            &cr.Number,
            &cr.CreatedAt,
            &cr.UpdatedAt); err != nil {
                return []internal.CineRoom{}, err
            }
            rooms = append(rooms, cr)
    }

    if err := rows.Err(); err != nil {
        return []internal.CineRoom{}, err
    }

    return rooms, nil 
}
     
func (r *roomRepository) GetRoomById(ctx context.Context, theaterId, roomId uint) (internal.CineRoom, error) {
    room := internal.CineRoom{Id: roomId}
    if err :=
        r.Conn.QueryRow(
            ctx,
            `SELECT number, created_at, updated_at
                FROM rooms WHERE
                theater_id = $1 AND id = $2 AND deleted_at IS NULL`, theaterId, roomId).Scan(&room.Number, &room.CreatedAt, &room.UpdatedAt); err != nil {
                    if err == pgx.ErrNoRows {
                        return internal.CineRoom{}, nil
                    }
                    return internal.CineRoom{}, err
    }
	return room, nil
}

func (r *roomRepository) ListAvailableRoomSeats(ctx context.Context, theaterId uint, room string) ([]internal.Seat, error) {
    rows, err :=
        r.Conn.Query(
            ctx,
            `SELECT s.row, s.place_number, cr.number AS rooms, s.available
                FROM seats AS s
                JOIN rooms AS cr ON s.cine_room = cr.number
                WHERE cr.theater_id = $1 AND cr.number = $2 AND deleted_at IS NULL`, theaterId, room)

    if err != nil {
        return []internal.Seat{}, err
    }
    
    defer rows.Close()

    seats := []internal.Seat{}
    for rows.Next() {
        s := internal.Seat{}
        if err := rows.Scan(
            &s.Row,
            &s.PlaceNumber,
            &s.CineRoom,
            &s.Available); err != nil {
            return []internal.Seat{}, err
        }
        seats = append(seats, s)
    }

    return seats, nil
}
