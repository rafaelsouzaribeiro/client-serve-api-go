/*
* O service irá atender as requisições externas que
* batem na nossa api, o usecase é a nossa regra de negócio
* e o repository é nosso adapter do banco de dados.
 */

package domain

import (
	"context"
)

type ApiResult struct {
	Bid string `json:"bid"`
}

// ProductService is a contract of http adapter layer
type PriceService interface {
	Get(ctx context.Context) error
}

// ProductUseCase is a contract of business rule layer
type PriceUseCase interface {
	Get() (*ApiResult, error)
}

// ProductRepository is a contract of database connection adapter
type PriceRepository interface {
	Insert(filename, content string) error
}
