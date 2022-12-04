package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	Surname       string `json:"surname"`
	Name          string `json:"name"`
	Placeofbirth  string `json:"placeofbirth"`
	Age           string `json:"age"`
	Weight        string `json:"weight"`
	Nationality   string `json:"nationality"`
	MaritalStatus string `json:"maritalstatus"`
	Address       string `json:"address"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Status        int    `json:"status"`
	Sex           string `json:"sex"`
}

type GetUser struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	Surname       string             `json:"surname"`
	Name          string             `json:"name"`
	Placeofbirth  string             `json:"placeofbirth"`
	Age           string             `json:"age"`
	Weight        string             `json:"weight"`
	Nationality   string             `json:"nationality"`
	MaritalStatus string             `json:"maritalstatus"`
	Address       string             `json:"address"`
	Email         string             `json:"email"`
	Phone         string             `json:"phone"`
	Status        int                `json:"status"`
	Sex           string             `json:"sex"`
}

type Users struct {
	users []User
}
