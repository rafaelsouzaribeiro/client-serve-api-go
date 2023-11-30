package pricerepository

import (
	"context"

	"github.com/rafaelsouzaribeiro/client-serve-api-go/server/core/domain"
)

func (r *Repository) Insert(cxt context.Context, values *domain.Price) error {
	sqlInsert := "INSERT INTO price(code,code_in,name,high,low,var_bid,pct_change,bid,ask,timestamp,create_date) VALUES (?,?,?,?,?,?,?,?,?,?,?)"

	stmt, err := r.db.PrepareContext(cxt, sqlInsert)

	if err != nil {
		return err
	}

	defer stmt.Close()

	v := &values.Usdbrl
	_, err = stmt.ExecContext(cxt, v.Code, v.Codein,
		v.Name, v.High, v.Low,
		v.VarBid, v.PctChange, v.Bid,
		v.Ask, v.Timestamp, v.CreateDate)

	if err != nil {
		return err
	}

	return nil
}
