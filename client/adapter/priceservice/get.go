package priceservice

import "context"

func (s *service) Get(ctx context.Context) error {
	currentQuote, err := s.usecase.Get()
	if err != nil {
		return err
	}

	err = s.repository.Insert("cotacao.txt", "Dólar:{"+currentQuote.Bid+"}")
	if err != nil {
		return err
	}

	return nil
}
