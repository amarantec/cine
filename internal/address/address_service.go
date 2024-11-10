package address

import (
	"context"
	"gitlab.com/amarantec/cine/internal"
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
	return s.addressRepository.InsertAddress(ctx, address)
}

func (s *addressService) GetAddress(ctx context.Context, id uint) (internal.Address, error) {
	return s.addressRepository.GetAddress(ctx, id)
}

func (s *addressService) UpdateAddress(ctx context.Context, address internal.Address) (bool, error) {
	return s.addressRepository.UpdateAddress(ctx, address)
}

func (s *addressService) DeleteAddress(ctx context.Context, id uint) (bool, error) {
	return s.addressRepository.DeleteAddress(ctx, id)
}

