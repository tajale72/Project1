package controller

import (
	"encoding/json"
	"time"

	"myapi/model"
)

func (s *Service) StoreFinancial(body []byte) (model.Finance, error) {
	var finance model.Finance
	json.Unmarshal(body, &finance)
	finance.Date = time.Now()
	total := -finance.Discoveracardbalance + finance.DebitCardbalance + finance.Stockmarket_buyingpower - finance.Rent - finance.Utilities - finance.Phonebill - finance.Carinsurance - finance.WalmartGrocery
	s.DB.StoreFinancial(finance, total)
	return finance, nil
}
