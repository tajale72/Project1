package controller

import (
	_ "github.com/lib/pq"

	"myapi/db"
	"myapi/model"
)

type Service struct {
	DB db.DatabaseInterface
}

type ControllerInterface interface {
	Hello() (string, error)
	VerifyToken(token string) (string, error)
	GenerateToken(body []byte) (string, error)
	StoreFinancial(body []byte) (model.Finance, error)
}

func (s *Service) Hello() (string, error) {
	return s.DB.Hello()
}
