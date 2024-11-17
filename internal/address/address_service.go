package address

import (
	"context"
	"gitlab.com/amarantec/cine/internal"
	"errors"
	"unicode/utf8"
    "unicode"
)	


type AddressService interface {
	InsertAddress(ctx context.Context, address internal.Address) (uint, error)
	GetAddress(ctx context.Context, id uint) (internal.Address, error)
	UpdateAddress(ctx context.Context, address internal.Address) (bool, error)	
	DeleteAddress(ctx context.Context, id uint) (bool, error)	
}

type addressService struct {
	addressRepository AddressRepository
}

func NewAddressService(repository AddressRepository) AddressService {
	return &addressService{addressRepository: repository}
}


func (s *addressService) InsertAddress(ctx context.Context, address internal.Address) (uint, error) {
    if valid, err := validateAddress(address); err != nil || !valid {
        return internal.ZERO, err
    }
	return s.addressRepository.InsertAddress(ctx, address)
}

func (s *addressService) GetAddress(ctx context.Context, id uint) (internal.Address, error) {
	return s.addressRepository.GetAddress(ctx, id)
}

func (s *addressService) UpdateAddress(ctx context.Context, address internal.Address) (bool, error) {
    if valid, err := validateAddress(address); err != nil || !valid {
        return false, err
    }
	return s.addressRepository.UpdateAddress(ctx, address)
}

func (s *addressService) DeleteAddress(ctx context.Context, id uint) (bool, error) {
	return s.addressRepository.DeleteAddress(ctx, id)
}

func validateAddress(a internal.Address) (bool, error) {
	if a.City == "" {
		return false, ErrAddressCityEmpty
	} else if utf8.RuneCountInString(a.City) < 2 || utf8.RuneCountInString(a.City) > 50 {
        return false, ErrAddressCityInvalidFormat
    }

	if a.Street == "" {
		return false, ErrAddressStreetEmpty
	} else if utf8.RuneCountInString(a.Street) < 2 || utf8.RuneCountInString(a.Street) > 50 {
        return false, ErrAddressStreetInvalidFormat
    }
	
	if a.ZIP == "" {
		return false, ErrAddressZipEmpty
	} else if utf8.RuneCountInString(a.ZIP) != 8 {
        return false, ErrAddressZipInvalidFormat
    } 

    for _, char := range a.ZIP {
        if !unicode.IsDigit(char) {
            return false, ErrAddressZipInvalidFormat
        }
    }

	if a.State == "" {
		return false, ErrAddressStateEmpty
	} else if utf8.RuneCountInString(a.State) != 2 {
        return false, ErrAddressStateInvalidFormat
    }

	if a.Country == "" {
		return false, ErrAddressCountryEmpty
	} else if utf8.RuneCountInString(a.Country) != 2 {
        return false, ErrAddressCountryInvalidFormat
    }

	return true, nil
}

var ErrAddressCityEmpty = errors.New("address city is empty")
var ErrAddressStreetEmpty = errors.New("address street is empty")
var ErrAddressZipEmpty = errors.New("address zip is empty")
var ErrAddressStateEmpty = errors.New("address state is empty")
var ErrAddressCountryEmpty = errors.New("address country is empty")

var ErrAddressCityInvalidFormat = errors.New("address city must be between 2-50 characters")
var ErrAddressStreetInvalidFormat = errors.New("address street must be between 2-50 characters")
var ErrAddressZipInvalidFormat = errors.New("address zip must contain 8 digits in range 0-9")
var ErrAddressStateInvalidFormat = errors.New("address state must contain 2 characters. ex: RS, SP")
var ErrAddressCountryInvalidFormat = errors.New("address country must contain 2 characters. ex: BR, US")
