package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"

	"myapi/model"
)

//InsertUser inserts data into the test database
func (s *Service) InsertAllUsers(users model.User) (*mongo.InsertOneResult, error) {

	collection := s.Mongoclient.Database("test").Collection("consultants")
	res, err := collection.InsertOne(context.Background(), users)
	if err != nil {
		return nil, err
	}
	return res, nil

}

//InsertUser inserts data into the test database
func (s *Service) GetAllUsers() ([]model.User, error) {
	var users []model.User
	collection := s.Mongoclient.Database("test").Collection("consultants")

	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("error from getting an object ", err)
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var user model.User
		err := cur.Decode(&user)
		if err != nil {
			log.Println("error from getting single data from cur", err)
		}
		users = append(users, user)
	}
	return users, nil

}

//InsertUser inserts data into the test database
func (s *Service) GetUserById(id string) (*model.User, error) {
	var user model.User
	collection := s.Mongoclient.Database("test").Collection("consultants")
	err := collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		log.Println("error from getting an object ", err)
		return nil, err
	}
	return &user, nil

}

//InsertUser inserts data into the test database
func (s *Service) DeleteUserById(id string) (*mongo.DeleteResult, error) {
	collection := s.Mongoclient.Database("test").Collection("consultants")
	res, err := collection.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		log.Println("error from getting an object ", err)
		return nil, err
	}
	return res, nil

}
