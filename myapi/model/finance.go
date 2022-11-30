package model

import "time"

type Finance struct {
	UserID                  string    `json:"user_id`
	Date                    time.Time `json:"date"`
	Discoveracardbalance    float64   `json:"discoveracardbalance"`
	DebitCardbalance        float64   `json:"debitcardbalance"`
	Stockmarket_buyingpower float64   `json:"stockmarket_buyingpower"`
	Rent                    float64   `json:"rent"`
	Utilities               float64   `json:"utilities"`
	Phonebill               float64   `json:phone_bill"`
	Carinsurance            float64   `json:"carinsurance"`
	WalmartGrocery          float64   `json:"walmart_grocery"`
	Networth                float64   `json: "netWorth"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Id    string `json:"id"`
}

type Users struct {
	users []User
}
