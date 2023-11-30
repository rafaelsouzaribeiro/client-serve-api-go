package di

import (
	"github.com/rafaelsouzaribeiro/client-serve-api-go/server/adapter/http/priceservice"
	"github.com/rafaelsouzaribeiro/client-serve-api-go/server/adapter/sqllite"
	"github.com/rafaelsouzaribeiro/client-serve-api-go/server/adapter/sqllite/pricerepository"
	"github.com/rafaelsouzaribeiro/client-serve-api-go/server/core/domain"
	"github.com/rafaelsouzaribeiro/client-serve-api-go/server/core/domain/usecase/pricesecase"
)

// ConfigProductDI return a ProductService abstraction with dependency injection configuration
func ConfigProductDI(conn sqllite.DataBaseInterface) domain.PriceService {
	priceRepository := pricerepository.New(conn)
	priceUseCase := pricesecase.New(priceRepository)
	priceService := priceservice.New(priceUseCase)

	return priceService
}
