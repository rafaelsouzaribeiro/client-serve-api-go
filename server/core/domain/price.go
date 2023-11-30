/*
* O service irá atender as requisições externas que
* batem na nossa api, o usecase é a nossa regra de negócio
* e o repository é nosso adapter do banco de dados.
 */

package domain

import (
	"context"
	"net/http"
)

type ApiResult struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

// ProductService is a contract of http adapter layer
type PriceService interface {
	Get(response http.ResponseWriter, request *http.Request)
}

type Price struct {
	Usdbrl ApiResult `json:"USDBRL"`
}

// ProductUseCase is a contract of business rule layer
type PriceUseCase interface {
	Get() (*Price, error)
	Insert(ctx context.Context, price *Price) error
}

// ProductRepository is a contract of database connection adapter
type PriceRepository interface {
	Insert(ctx context.Context, result *Price) error
}
