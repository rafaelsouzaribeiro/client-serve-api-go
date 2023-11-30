package pricesecase

import (
	"context"
	"time"

	"github.com/rafaelsouzaribeiro/client-serve-api-go/server/core/domain"
)

func (usecase usecase) Insert(ctx context.Context, data *domain.Price) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err := usecase.repository.Insert(ctx, data)

	if err != nil {
		return err
	}

	return nil
}
