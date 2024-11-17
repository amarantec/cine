package theater

import (
	"context"
	"errors"
	"gitlab.com/amarantec/cine/internal"
	"unicode/utf8"
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
	return s.theaterRepository.GetTheaterById(ctx, id)
}

func (s *theaterService) AddTheater(ctx context.Context, theater internal.Theater) (uint, error) {
	if valid, err := validateTheater(theater); err != nil || !valid {
		return internal.ZERO, err
	}

	return s.theaterRepository.AddTheater(ctx, theater)
}

func (s *theaterService) UpdateTheater(ctx context.Context, theater internal.Theater) (bool, error) {
	if valid, err := validateTheater(theater); err != nil || !valid {
		return false, err
	}

	return s.theaterRepository.UpdateTheater(ctx, theater)
}

func (s *theaterService) DeleteTheater(ctx context.Context, id uint) (bool, error) {
	return s.theaterRepository.DeleteTheater(ctx, id)
}

func validateTheater(t internal.Theater) (bool, error) {
	if t.Name == "" {
		return false, ErrTheaterNameEmpty
	}	
	
	if t.AddressId == nil {
		return false, ErrAddressIdEmpty
	}

	if utf8.RuneCountInString(t.Name) < 2 && utf8.RuneCountInString(t.Name) > 50 {
		return false, ErrTheaterNameInvalidFormat
	}

	return true, nil
}

var ErrTheaterIdEmpty = errors.New("theater id is should be greater than 0")
var ErrTheaterNameEmpty = errors.New("theater name is empty")
var ErrAddressIdEmpty = errors.New("theater address id is empty")
var ErrTheaterNameInvalidFormat = errors.New("theater name must be beetwen 2-50 characters")
