package priceservice

import (
	"github.com/rafaelsouzaribeiro/client-serve-api-go/server/core/domain"
)

type service struct {
	usecase domain.PriceUseCase
}

// New returns contract implementation of PriceService
func New(usecase domain.PriceUseCase) domain.PriceService {
	return &service{
		usecase: usecase,
	}
}
