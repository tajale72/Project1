package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"

	"interview/internal/model"
)

// var _ Businesslogic = &Service{}

// type Businesslogic interface {
// 	InsertUser(body []byte) (*mongo.InsertOneResult, error)
// 	GetUser() (model.User, error)
// 	Mongo() (*mongo.Client, error)
// }

type Service struct {
	dbclient *mongo.Client
}

// //InsertUser inserts data into the test database
// func (s *Service) InsertUser(user model.User) (*mongo.InsertOneResult, error) {
// 	sort.Strings(user.Arr)
// 	collection := s.dbclient.Database("test").Collection("user")
// 	res, err := collection.InsertOne(context.Background(), user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil

// }

//GetUser get's data from the database
func (s *Service) GetUser(name string) ([]model.User, error) {
	var user model.User
	var a []model.User
	collection := s.dbclient.Database("test").Collection("user")
	err := collection.FindOne(context.Background(), bson.M{"name": name}).Decode(&user)
	a = append(a, user)
	if err != nil {
		log.Println("error getting response from collection client", err)
		return nil, err
	}
	return a, nil
}

func (s *Service) Mongo() (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	s.dbclient = client
	return client, nil
}
