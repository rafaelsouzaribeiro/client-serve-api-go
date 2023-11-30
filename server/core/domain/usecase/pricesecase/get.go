package pricesecase

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/rafaelsouzaribeiro/client-serve-api-go/server/core/domain"
)

func (u *usecase) Get() (*domain.Price, error) {
	cxt, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)

	defer cancel()

	request, err := http.NewRequestWithContext(cxt, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	api := domain.Price{}
	err = json.Unmarshal(body, &api)

	return &api, nil
}
