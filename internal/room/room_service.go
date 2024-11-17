package room

import (
    "context"
    "gitlab.com/amarantec/cine/internal"
	"errors"
	"unicode"
)

type RoomService interface {
	ListRooms(ctx context.Context, theaterId uint) ([]internal.CineRoom, error) 
	GetRoomById(ctx context.Context, theaterId, roomId uint) (internal.CineRoom, error)
	ListAvailableRoomSeats(ctx context.Context, theaterId uint, room string) ([]internal.Seat, error)
}

type roomService struct {
    roomRepository RoomRepository
}

func NewRoomService(repository RoomRepository) RoomService {
    return &roomService{roomRepository: repository}
}

func (s *roomService) ListRooms(ctx context.Context, theaterId uint) ([]internal.CineRoom, error) {
	return s.roomRepository.ListRooms(ctx, theaterId)	
}

func (s *roomService) GetRoomById(ctx context.Context, theaterId, roomId uint) (internal.CineRoom, error) {
	return s.roomRepository.GetRoomById(ctx, theaterId, roomId)
}
	 
func (s *roomService) ListAvailableRoomSeats(ctx context.Context, theaterId uint, room string) ([]internal.Seat, error) {
	if valid, err := validRoomNumber(room); err != nil || !valid {
		return []internal.Seat{}, errors.New("could not list available seats, error: " + err.Error())
	}	
	
	return s.roomRepository.ListAvailableRoomSeats(ctx, theaterId, room)
}

func validRoomNumber(r string) (bool, error) {
	if r == "" {
		return false, ErrRoomNumberEmpty
	}

	for _, char := range r {
		if !unicode.IsDigit(char) {
			return false, ErrRoomNumberInvalidFormat
		}
	}
	return true, nil
}
	
var ErrTheaterIdEmpty = errors.New("theater id is empty")
var ErrRoomNumberEmpty = errors.New("room number is empty")
var ErrRoomNumberInvalidFormat = errors.New("room number must be beetwen range at 0-9")
