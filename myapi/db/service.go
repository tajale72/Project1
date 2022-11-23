package db

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"

	"myapi/model"
)

type Service struct {
	Mongoclient *mongo.Client
	Db          *sql.DB
	Name        string
}

type DatabaseInterface interface {
	Hello() (string, error)
	StoreFinancial(finance model.Finance, total float64)
}

func (s *Service) Hello() (string, error) {
	return "hello", nil
}
