package priceusecase

import (
	"github.com/rafaelsouzaribeiro/client-serve-api-go/client/core/domain"
)

type usecase struct {
}

func New() domain.PriceUseCase {
	return &usecase{}
}
