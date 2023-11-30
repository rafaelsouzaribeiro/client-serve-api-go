package di

import (
	"github.com/rafaelsouzaribeiro/client-serve-api-go/client/adapter/priceservice"
	"github.com/rafaelsouzaribeiro/client-serve-api-go/client/core/domain"
	"github.com/rafaelsouzaribeiro/client-serve-api-go/client/file/pricerepository"

	"github.com/rafaelsouzaribeiro/client-serve-api-go/client/core/domain/usecase/priceusecase"
)

func ConfigPriceDI() domain.PriceService {
	repository := pricerepository.New()
	usecase := priceusecase.New()
	service := priceservice.New(usecase, repository)
	return service
}
