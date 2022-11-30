package controller

import (
	"encoding/json"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"myapi/model"
)

func (s *Service) InsertAllUsers(body []byte) (*mongo.InsertOneResult, error) {
	var user model.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}
	log.Println(len(user.Id))
	if len(user.Id) == 0 {
		return nil, errors.New("error missing id")
	} else {
		log.Println("why am i here")
		return s.DB.InsertAllUsers(user)
	}

}

func (s *Service) GetAllUsers() ([]model.User, error) {
	return s.DB.GetAllUsers()
}

func (s *Service) GetUserById(id string) (*model.User, error) {
	return s.DB.GetUserById(id)
}

func (s *Service) DeleteUserById(id string) (*mongo.DeleteResult, error) {
	return s.DB.DeleteUserById(id)
}
