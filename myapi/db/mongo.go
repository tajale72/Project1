package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"myapi/model"
)

func Mongo() (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	//s.dbclient = client
	return client, nil
}

//InsertUser inserts data into the test database
func (s *Service) InsertUser(finance model.Finance, total float64) (*mongo.InsertOneResult, error) {
	//sort.Strings(user.Arr)
	log.Println("i am in InsertUser")
	finance.Networth = total
	collection := s.Mongoclient.Database("test").Collection("finance")
	res, err := collection.InsertOne(context.Background(), finance)
	if err != nil {
		return nil, err
	}
	return res, nil

}
