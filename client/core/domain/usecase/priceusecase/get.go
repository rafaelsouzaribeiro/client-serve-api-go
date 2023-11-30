package priceusecase

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rafaelsouzaribeiro/client-serve-api-go/client/core/domain"
)

func (u *usecase) Get() (*domain.ApiResult, error) {
	cxt, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)

	defer cancel()

	request, err := http.NewRequestWithContext(cxt, "GET", "http://localhost:8080/cotacao", nil)

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	// body, err := io.ReadAll(res.Body)

	// if err != nil {
	// 	return nil, err
	// }

	defer res.Body.Close()

	var apiResult domain.ApiResult

	err = json.NewDecoder(res.Body).Decode(&apiResult)

	if err != nil {
		return nil, err
	}

	return &apiResult, nil
}
