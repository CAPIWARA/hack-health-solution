package users

import (
	"time"
	"github.com/mitchellh/mapstructure"
	"errors"
	"github.com/kr/pretty"
)

type User struct {
	Id       string    `json:"-" bson:"_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	UpdateAt time.Time `json:"updateAt"`
}

func GetUser(id string) (*User, error) {
	var user *User
	res, err := db.MongoRepoBuilder(UserCollection).FindById(id)

	if err != nil {
		pretty.Log("error: ", err)
		return nil, err
	}
	if err = mapstructure.Decode(res, &user); err != nil {
		pretty.Log("error: ", err)
		return nil, err
	}
	return user, nil
}

func FindByEmail(email string) (*User, error) {
	var data *User
	res, err := db.MongoRepoBuilder(UserCollection).FindByQuery("email", email)
	if err != nil {
		pretty.Log("error: ", err)
		return nil, err
	}
	if res == nil {
		pretty.Log("error: ", err)
		return nil, errors.New("user not found")
	}
	if err := mapstructure.Decode(res, &data); err != nil {
		pretty.Log("error: ", err)
		return nil, err
	}
	return data, nil
}
