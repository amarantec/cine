package theater

import (
	"context"
	"errors"
	"gitlab.com/amarantec/cine/internal"
)

type TheaterService interface {
	ListTheaters(ctx context.Context) ([]internal.Theater, error)
	GetTheaterById(ctx context.Context, id uint) (internal.Theater, error)
	AddTheater(ctx context.Context, theater internal.Theater) (uint, error)
	UpdateTheater(ctx context.Context, theater internal.Theater) (bool, error)
	DeleteTheater(ctx context.Context, id uint) (bool, error)
}

type theaterService struct {
	theaterRepository TheaterRepository
}

func NewTheaterService(repository TheaterRepository) TheaterService {
	return &theaterService{theaterRepository: repository}
}

func (s *theaterService) ListTheaters(ctx context.Context) ([]internal.Theater, error) {
	return s.theaterRepository.ListTheaters(ctx)
}

func (s *theaterService) GetTheaterById(ctx context.Context, id uint) (internal.Theater, error) {
	if id <= 0 {
		return internal.Theater{}, ErrTheaterIdEmpty
	}

	return s.theaterRepository.GetTheaterById(ctx, id)
}

func (s *theaterService) AddTheater(ctx context.Context, theater internal.Theater) (uint, error) {
	return s.theaterRepository.AddTheater(ctx, theater)
}

func (s *theaterService) UpdateTheater(ctx context.Context, theater internal.Theater) (bool, error) {
	return s.theaterRepository.UpdateTheater(ctx, theater)
}

func (s *theaterService) DeleteTheater(ctx context.Context, id uint) (bool, error) {
	return s.theaterRepository.DeleteTheater(ctx, id)
}

var ErrTheaterIdEmpty = errors.New("theater id is should be greater than 0")
