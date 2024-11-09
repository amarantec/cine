package theater

import (
	"context"

	"gitlab.com/amarantec/cine/internal"
)

type TheaterService interface {
	ListTheaters(ctx context.Context) ([]internal.Theater, error)
	GetTheaterById(ctx context.Context, id uint) (internal.Theater, error)
	AddTheater(ctx context.Context, theater internal.Theater) (uint, error)
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
	return s.theaterRepository.GetTheaterById(ctx, id)
}

func (s *theaterService) AddTheater(ctx context.Context, theater internal.Theater) (uint, error) {
	return s.theaterRepository.AddTheater(ctx, theater)
}
