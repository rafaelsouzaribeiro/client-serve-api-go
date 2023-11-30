package priceservice

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (s *service) logError(response http.ResponseWriter, err error) {
	log.Printf("Erro [%s]: %v", time.Now().Format(time.DateTime), err)
	response.WriteHeader(http.StatusInternalServerError)
	response.Write([]byte(fmt.Sprintf("Ocorreu um erro: %v", err)))

}

func (s *service) Get(response http.ResponseWriter, request *http.Request) {
	resp, err := s.usecase.Get()

	if err != nil {
		s.logError(response, err)
		return
	}

	err = s.usecase.Insert(nil, resp)
	if err != nil {
		s.logError(response, err)
		return
	}

	bidResponse := struct {
		Bid string `json:"bid"`
	}{
		Bid: resp.Usdbrl.Bid,
	}

	jsonData, err := json.Marshal(bidResponse)

	if err != nil {
		panic(err)
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonData)

}
