package model

import "time"

type Finance struct {
	Date                    time.Time `json:"date"`
	Discoveracardbalance    float64
	DebitCardbalance        float64
	Stockmarket_buyingpower float64
	Rent                    float64
	Utilities               float64
	Phonebill               float64
	Carinsurance            float64
	WalmartGrocery          float64
	Networth                float64
}
